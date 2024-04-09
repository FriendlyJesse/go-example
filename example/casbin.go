package example

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

func ExecCasbin() {

	a, _ := gormadapter.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3306)/test", true)
	casbin.NewEnforcer("./model.conf", a)

	// // 加载策略
	// e.LoadPolicy()

	// // 测试 request
	// sub := "alice" // the user that wants to access a resource.
	// obj := "data1" // the resource that is going to be accessed.
	// act := "read"  // the operation that the user performs on the resource.
	// ok, err := e.Enforce(sub, obj, act)

	// if err != nil {
	// 	fmt.Println("err: ", err)
	// }

	// if ok {
	// 	fmt.Println("通过")
	// } else {
	// 	fmt.Println("未通过")
	// }
}
