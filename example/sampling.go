package example

import (
	"fmt"
	"math/rand"
	"time"
)

func Sampling() {

	var num int
	var products = map[int]int{}
	var sample []int
	// 设置随机数种子
	var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println("请输入检测产品数量")
	fmt.Scanln(&num)

	// 生成产品编号
	for i := 1; i <= num; i++ {
		products[i] = i
	}
	fmt.Printf("产品编号：%v\n", products)

	// 在产品数量的 20% ~ 70% 范围选择样品数量
	// 假设：num = 10, 那么就是：0~5 + 2
	// 假设 num = 0，那么0+2就是 20%; 假设 num=5，那么5+2就是 70%
	var runs = rnd.Intn(num/2) + num/4
	fmt.Printf("抽查率：%.2f%%\n", float64(runs)/float64(num)*100)

	for i := 1; i <= runs; i++ {
		var n = rnd.Intn(num) + 1
		_, ok := products[n]
		// 如果存在则加入样品
		if ok {
			sample = append(sample, n)
			delete(products, n)
		} else {
			i--
		}
	}
	fmt.Println("样品：", sample)

	var qualified = 0
	for _, v := range sample {
		var probability = rnd.Intn(100) + 1

		if probability > 50 {
			fmt.Printf("产品编号：%v ，检测合格\n", v)
			qualified++
		} else {
			fmt.Printf("产品编号：%v ，检测不合格\n", v)
		}
	}

	var rate = float64(qualified) / float64(len(sample)) * 100
	fmt.Printf("合格率：%.2f%%\n", rate)
}
