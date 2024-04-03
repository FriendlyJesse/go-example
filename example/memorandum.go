package example

import (
	"fmt"
	"sort"
	"time"
)

type memorandum struct {
	Date  int64
	Event string
}

type memorandumSort []memorandum

func (m memorandumSort) Len() int {
	return len(m)
}

func (m memorandumSort) Less(i, j int) bool {
	if m[i].Date < m[j].Date {
		return true
	} else if m[i].Date > m[j].Date {
		return false
	} else {
		return m[i].Event < m[j].Event
	}
}

func (m memorandumSort) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func ExecMemo() {
	var now = time.Now().Unix()

	var m1 = memorandum{Date: now, Event: "学习 Golang"}
	var m2 = memorandum{Date: now + 7250, Event: "继续学习 Golang"}
	var m3 = memorandum{Date: now + 9070, Event: "晚了，洗洗睡吧"}
	var m4 = memorandum{Date: now + 3460, Event: "休息吧，顺便吃顿饭"}
	var m = memorandumSort{m1, m2, m3, m4}
	sort.Stable(m)
	for _, v := range m {
		var t = time.Unix(v.Date, 0).Format("2006-01-02 15:04:05")
		fmt.Printf("备忘时间：%v，你要做：%v\n", t, v.Event)
	}
}
