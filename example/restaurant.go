package example

import (
	"fmt"
	"sync"
	"time"
)

// 同步等待组
var wg sync.WaitGroup

func producer(intChan chan int, exitChan chan bool) {
	// 往通道 int 写入数据，每两秒生成一道菜
	for i := 0; i < cap(intChan); i++ {
		fmt.Printf("厨师完成菜式%v的制作\n", i)
		intChan <- i
		time.Sleep(2 * time.Second)
	}
	// 往通道 exitChan 写入数据
	exitChan <- true
	fmt.Println("厨师完成所有菜式")
	wg.Done()
}

func consumer(intChan chan int, exitChan chan bool) {
fors:
	for {
		select {
		case v, ok := <-intChan:
			if ok { // 如果能够读取数据，说明完成了一次生成
				fmt.Printf("顾客吃完了菜式：%v\n", v)
			}
		case v := <-exitChan:
			if v { // 如果读取到 true，说明该下班了
				fmt.Println("厨师下班，店铺不经营！")
				break fors
			}

		default: // 如果前者都没有动静，说明顾客需要等待
			fmt.Println("顾客等待中...")
			time.Sleep(3 * time.Second)
		}
	}
	wg.Done()
}

func ExecRestaurant() {
	wg.Add(2)

	var intChan = make(chan int, 3)
	var exitChan = make(chan bool)

	go producer(intChan, exitChan)
	go consumer(intChan, exitChan)

	wg.Wait()
	fmt.Println("main 解除阻塞")
}
