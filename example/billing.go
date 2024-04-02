package example

import "fmt"

func ExecBilling() {
	var distance, cost float64

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕捉异常：", err)
		} else {
			fmt.Printf("当前里程数：%v，车费：%v\n", distance, cost)
		}
	}()

	fmt.Println("输入里程数：")
	fmt.Scanln(&distance)

	if distance < 0 {
		panic("公里数小于0，无法计算车费！")
	} else if distance <= 3 {
		cost = 13.0
	} else if distance <= 10 {
		cost = 13.0 + (distance-3)*2.3
	} else {
		cost = 13.0 + (10-3)*2.3 + (distance-10)*3.2
	}

}
