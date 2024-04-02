package example

import (
	"encoding/json"
	"fmt"
	"os"
)

type personinfo struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

func ExecWriteJson() {
	// 使用OpenFile()打开文件,设置O_TRUNC模式,每次写入将覆盖原有数据
	// 如果不想为OpenFile()设置参数,则可以用Create()代替,实现效果一样
	f, _ := os.OpenFile("output.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	// 创建PersonInfo类型的切片
	p := []personinfo{{"David", 30}, {"Lee", 27}}
	// 实例化结构体Encoder,实现数据写入
	encoder := json.NewEncoder(f)
	// 将变量p的数据写入JSON文件
	// 数据写入必须使用文件内容覆盖,即设置os.0_TRUNC模式,否则导致内容错乱
	err := encoder.Encode(p)
	// 如果err不为空值nil,则说明写入错误
	if err != nil {
		fmt.Printf("JSON写入失败；%v\n", err.Error())
	} else {
		fmt.Printf("JSON写入成功\n")
	}
}

func ExecReadJson() {
	//使用OpenFile()打开文件,设置O_CREATE模式,若文件不存在则创建
	//如果不想为OpenFile()设置参数,则可以用Open(）代替,实现效果一样
	f, _ := os.OpenFile("output.json", os.O_RDWR|os.O_CREATE, 0755)
	// 定义结构体类型的切片
	var person []personinfo
	// 实例化结构体Decoder,实现数据读取
	data := json.NewDecoder(f)
	// 将已读取的数据加载到切片person
	err := data.Decode(&person)
	// 如果err不为空值nil,则说明读取错误
	if err != nil {
		fmt.Printf("JSON读取失败：%v\n", err.Error())
	} else {
		fmt.Printf("JSON读取成功：%v\n", person)
		// 关闭文件
		f.Close()
	}
}
