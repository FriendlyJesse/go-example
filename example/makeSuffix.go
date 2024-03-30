package example

import (
	"fmt"
	"path"
)

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if path.Ext(name) == "" {
			return name + suffix
		} else {
			return "文件已有后缀名：" + name
		}
	}
}

func ExecMakeSuffixFunc() {
	var jpgFunc = makeSuffixFunc(".jpg")
	var txtFunc = makeSuffixFunc(".txt")

	fmt.Println(jpgFunc("test.png")) // 文件已有后缀名：test.png
	fmt.Println(txtFunc("test"))     // test.txt
}
