package example

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func ExecWriteCsv() {
	// OpenFile(）创建或打开文件，设置读写模式
	// 如果设置O_APPEND模式，则实现文件续写功能
	// 如果设置OTRUNC模式，则新数据覆盖文件原有数据
	var nfs, _ = os.OpenFile("input.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	csvWriter := csv.NewWriter(nfs)

	// 设置分隔符，默认为逗号
	// csvWriter.Comma = ','
	// 设置换行符
	// csvWriter.UseCRLF = true

	var row = []string{"1", "2", "3", "4"}
	var err = csvWriter.Write(row)
	if err != nil {
		fmt.Println("无法写入，错误信息：", err)
	}

	// 一次性写入多行数据
	var newContent [][]string
	newContent = append(newContent, []string{"11", "12", "13", "14"})
	newContent = append(newContent, []string{"21", "22", "23", "24"})
	csvWriter.WriteAll(newContent)

	// 将数据写入文件
	csvWriter.Flush()
	nfs.Close()
}

func ExecReadCsv() {
	// OpenFile(）创建或打开文件，设置读写模式
	// O_RDWR 支持读写模式
	// os.O_CREATE 文件不存在时创建
	var fs, _ = os.OpenFile("input.csv", os.O_RDWR|os.O_CREATE, 0666)
	var csvReader = csv.NewReader(fs)

	// 一行一行地读取文件，常用于大文件
	for {
		row, err := csvReader.Read()
		if err == io.EOF || err != nil {
			break
		}
		fmt.Printf("逐行读取csv内容：%v，数据类型：%T\n", row, row)
	}
	fs.Close()

	// 一次性读取文件所有内容，常用于小文件
	var fs1, _ = os.OpenFile("input.csv", os.O_RDWR|os.O_CREATE, 0666)
	var csvReader1 = csv.NewReader(fs1)
	// 读取文件所有内容
	content, err := csvReader1.ReadAll()
	if err != nil {
		fmt.Println("读取失败：", err)
	}
	for _, v := range content {
		fmt.Printf("读取csv内容：%v，数据类型：%T\n", v, v)
	}
	fs1.Close()
}
