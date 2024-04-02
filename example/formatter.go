package example

import (
	"fmt"
	"reflect"
)

type person struct {
	Name string `key:"name"`
	Age  int    `key:"age"`
}

func ExecFormatter() {
	var result = map[string]any{}
	var p = person{Name: "张三", Age: 18}

	var pRT = reflect.TypeOf(p)
	var pRV = reflect.ValueOf(p)
	for i := 0; i < pRT.NumField(); i++ {
		var pInfo = pRT.Field(i)
		var pRVInfo = pRV.Field(i)
		var tag = pInfo.Tag
		var tagKey = tag.Get("key")
		result[tagKey] = pRVInfo
	}
	for k, v := range result {
		fmt.Println(k, v)
	}
}
