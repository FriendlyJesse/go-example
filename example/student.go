package example

import (
	"encoding/json"
	"fmt"
	"os"
)

type student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade string `json:"grade"`
	Major string `json:"major"`
}

func dataProces(style string, s ...student) []student {
	var students []student
	if style == "read" {
		f1, _ := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE, 0755)
		// 实例化结构体 Decoder, 实现数据读取
		var data = json.NewDecoder(f1)
		// 将已读取的数据加载到切片
		data.Decode(&students)
		f1.Close()
	} else if style == "write" {
		f2, _ := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		// 实例化结构体，实现数据写入
		var encoder = json.NewEncoder(f2)
		// 数据写入必须使用文件内容覆盖，即设置 os.O_TRUNC模式
		encoder.Encode(s)
		f2.Close()
	}
	return students
}

func ExecStu() {
	var s int
	for {
		fmt.Println("欢迎来到学生信息管理系统")
		fmt.Println("查询请按1，新增请按2，删除请按3，退出请按4")
		fmt.Scanln(&s)

		if s == 1 {
			data := dataProces("read")
			if len(data) == 0 {
				fmt.Println("数据为空！")
				continue
			}
			var qs int
			fmt.Println("查询全部请按1，查询某个学生请按2")
			fmt.Scanln(&qs)
			if qs == 1 {
				for _, v := range data {
					fmt.Printf("学号：%v\n", v.ID)
					fmt.Printf("姓名：%v\n", v.Name)
					fmt.Printf("年龄：%v\n", v.Age)
					fmt.Printf("年级：%v\n", v.Grade)
					fmt.Printf("专业：%v\n", v.Major)
				}
			} else if qs == 2 {
				var id int
				fmt.Println("请输入学号查询")
				fmt.Scanln(&id)
				for _, v := range data {
					if v.ID == id {
						fmt.Printf("学号：%v\n", v.ID)
						fmt.Printf("姓名：%v\n", v.Name)
						fmt.Printf("年龄：%v\n", v.Age)
						fmt.Printf("年级：%v\n", v.Grade)
						fmt.Printf("专业：%v\n", v.Major)
					}
				}
			}
		} else if s == 2 {
			var id, age int
			var name, grade, major string
			fmt.Println("请输入学号")
			fmt.Scanln(&id)
			fmt.Println("请输入姓名")
			fmt.Scanln(&name)
			fmt.Println("请输入年龄")
			fmt.Scanln(&age)
			fmt.Println("请输入年级")
			fmt.Scanln(&grade)
			fmt.Println("请输入专业")
			fmt.Scanln(&major)
			// 读取 json 文件获取学生信息
			var data = dataProces("read")
			var stu = student{
				ID:    id,
				Name:  name,
				Age:   age,
				Grade: grade,
				Major: major,
			}
			data = append(data, stu)
			dataProces("write", data...)
		} else if s == 3 {
			var id int
			var newData []student
			// 读取 json 文件获取学生信息
			var data = dataProces("read")
			fmt.Println("请输入学号删除学生信息")
			fmt.Scanln(&id)
			fmt.Printf("删除前的学生信息：%v\n", data)
			for _, v := range data {
				if v.ID != id {
					newData = append(newData, v)
				}
			}
			dataProces("write", newData...)
			fmt.Println("删除后的学生信息：", newData)
		} else if s == 4 {
			break
		}
	}
}
