package example

import "fmt"

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func fibonacciSum(arr []int, n int) int {
	return (2 * arr[n]) + (arr[n-1]) - 1
}

func ExecFibonacci() {
	var arr []int

	// 拿 1~10 的数列
	for i := 1; i <= 50; i++ {
		arr = append(arr, fibonacci(i))
	}

	// 数列求和
	var sum = fibonacciSum(arr, len(arr)-1)
	fmt.Println("数列之和：", sum)
}
