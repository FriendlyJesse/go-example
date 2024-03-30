package example

import (
	"fmt"
	"math/rand"
	"time"
)

func Loterry() {
	// 获取输入的号码
	var myNum int
	var myNums []int

	for i := 0; i < 7; i++ {
		fmt.Printf("请输入第%v位号码：\n", i+1)
		fmt.Scanln(&myNum)
		myNums = append(myNums, myNum)
	}
	fmt.Println("你选到的号码分别为：", myNums)

	// 获取随机号码
	var result []*int
	var status bool
	// 设置种子值
	var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		var num int
		status = false
		// 10 以内的随机数
		num = rnd.Intn(10) + 1

		// 判断是否有相同的号码
		for _, v := range result {
			if *v == num {
				status = true
			}
		}

		// 排除相同的号码
		if !status {
			result = append(result, &num)
		}
		if len(result) == 7 {
			break
		}
	}

	for i, v := range result {
		fmt.Printf("第%v位号码为：%v\n", i+1, *v)
	}

	// 找并集
	for _, v := range result {
		for _, j := range myNums {
			if *v == j {
				fmt.Printf("号码%v选中了\n", j)
			}
		}
	}
}
