package example

import "fmt"

type Person struct {
	Name    string
	Age     int
	Address string
}

// 普通查找数据
func findData(persons []*Person, name string, age int) {
	for _, data := range persons {
		if data.Name == name && data.Age == age {
			fmt.Println(data)
			return
		}
	}
	fmt.Println("没有找到对应的数据")
}

// 多键索引查询数据
type queryKey struct {
	Name string
	Age  int
}

// 定义集合， key 为结构体 queryKey, value 为 person 的指针
var mapper = map[queryKey]*Person{}

// 定义函数，建立多条件查询
func buildIndex(persons []*Person) {
	for _, p := range persons {
		var key = queryKey{
			Name: p.Name,
			Age:  p.Age,
		}
		mapper[key] = p
	}
	fmt.Printf("集合 mapper 是数据：%+v\n", mapper)
}

// 定义函数，用于查询数据
func queryData(name string, age int) {
	// 实例化结构体 queryKey, 作为查询条件
	var key = queryKey{Name: name, Age: age}
	// 从集合 mapper 中查询数据
	result, ok := mapper[key]
	if ok {
		fmt.Printf("查询结果：%+v\n", result)
	} else {
		fmt.Println("没有找到对应的数据")
	}
}

func ExecFindData() {
	var list = []*Person{
		{
			Name:    "Lily",
			Age:     23,
			Address: "CN",
		},
		{
			Name:    "Tom",
			Age:     25,
			Address: "CN",
		},
		{
			Name:    "Jesse",
			Age:     25,
			Address: "CN",
		},
	}
	// 普通查询数据
	findData(list, "Jesse", 25)

	// 多键索引查询数据
	buildIndex(list)
	queryData("Jesse", 25)
}

/*
多键索引查询的实现过程：
1. 定义两个结构体和一个集合，查询条件结构 K 作为 map 的键，数据内容结构 V 作为集合的值，集合执行数据查询。
2. 使用查询条件实例化结构体 Q，并与集合的键 K 进行匹配，找出对应的 V。从而完成整个查询过程。
*/
