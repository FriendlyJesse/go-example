package example

import "fmt"

func Calc() {
	var action string
	var d1, d2 float64
	var opt string
	var result any

	for {
		fmt.Println("请输入选择，按1计算，按2退出：")
		fmt.Scanln(&action)

		if action == "1" {
			for i := 0; i < 2; i++ {
				fmt.Println("请输入数字：")
				if i == 0 {
					fmt.Scanln(&d1)
				} else {
					fmt.Scanln(&d2)
				}
			}
		} else { // 退出无限循环
			break
		}

		fmt.Println("请输入运算法则，可选择+-*/：")
		fmt.Scanln(&opt)

		switch opt {
		case "+":
			result = d1 + d2
		case "-":
			result = d1 - d2
		case "*":
			result = d1 * d2
		case "/":
			if d2 != 0.0 {
				result = d1 / d2
			} else {
				result = "除数为零无法计算"
			}
		}
		fmt.Printf("%v %v %v = %v\n", d1, opt, d2, result)
	}
}
