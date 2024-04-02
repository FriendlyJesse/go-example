package example

import (
	"fmt"
	"sort"
)

type sortPerson struct {
	name string
	age  int
}

type personOnList []sortPerson

// 自定义 sort.Interface 的 Len()
func (list personOnList) Len() int {
	return len(list)
}

// 排序规则：首先按年龄排序（从小到大）
// 年龄相同时按姓名进行排序（按字符的自然排序）
// 自定义 sort.Interface 的 Less()
func (list personOnList) Less(i, j int) bool {
	if list[i].age < list[j].age {
		return true
	} else if list[i].age > list[j].age {
		return false
	} else {
		return list[i].name < list[j].name
	}
}

// 自定义 sort.Interface 的 Swap()
func (list personOnList) Swap(i, j int) {
	var temp sortPerson = list[i]
	list[i] = list[j]
	list[j] = temp
}

func ExecSort() {

	// []int 排序
	var nums = []int{2, 31, 5, 6, 3}
	sort.Ints(nums)
	fmt.Println("[]int顺序排序：", nums) // []int顺序排序： [2 3 5 6 31]
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Println("[]int倒序排序：", nums) // []int倒序排序： [31 6 5 3 2]

	// []float64 排序
	var floats = []float64{2.2, 6.6, -5.3, 6.66, 3.12}
	sort.Float64s(floats)
	fmt.Println("[]float64顺序排序：", floats) // []float64顺序排序： [-5.3 2.2 3.12 6.6 6.66]
	sort.Sort(sort.Reverse(sort.Float64Slice(floats)))
	fmt.Println("[]float64倒序排序：", floats) // []float64倒序排序： [6.66 6.6 3.12 2.2 -5.3]

	// []string 排序
	var strings = []string{"abc", "123", "kk", "Jordan", "Ko", "DD"}
	sort.Strings(strings)
	fmt.Println("[]string顺序排序：", strings) // []string顺序排序： [123 DD Jordan Ko abc kk]
	sort.Sort(sort.Reverse(sort.StringSlice(strings)))
	fmt.Println("[]string倒序排序：", strings) // []string倒序排序： [kk abc Ko Jordan DD 123]

	// 自定义排序
	var p1 = sortPerson{name: "Tom", age: 19}
	var p2 = sortPerson{name: "Hanks", age: 19}
	var p3 = sortPerson{name: "Amy", age: 19}
	var p4 = sortPerson{name: "Tom", age: 20}
	var p5 = sortPerson{name: "Jogn", age: 21}
	var p6 = sortPerson{name: "Mike", age: 23}
	var pList = personOnList([]sortPerson{p1, p2, p3, p4, p5, p6})
	sort.Sort(pList)
	fmt.Println("自定义排序：", pList) // 自定义排序： [{Amy 19} {Hanks 19} {Tom 19} {Tom 20} {Jogn 21} {Mike 23}]
	// Stable() 比 Sort() 稳定
	sort.Stable(pList)
	fmt.Println("自定义排序：", pList) // 自定义排序： [{Amy 19} {Hanks 19} {Tom 19} {Tom 20} {Jogn 21} {Mike 23}]

	// 二分查找法
	var index = sort.Search(len(pList), func(i int) bool {
		return pList[i].name == "Tom" && pList[i].age == 20
	})
	fmt.Printf("查找索引位置：%v，查找结果：%v\n", index, pList[index]) // 查找索引位置：3，查找结果：{Tom 20}
}
