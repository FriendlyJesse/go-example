package example

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type userinfo struct {
	ID      int
	Name    string
	Age     int
	Created string
}

func ExecMysql() {
	fmt.Println("mysql 启动")
	var dataSourceName = "root:123456@(127.0.0.1:3306)/test"
	var db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println(err)
		return
	}

	var sqlTable = `CREATE TABLE IF NOT EXISTS userinfo (
		id int PRIMARY KEY AUTO_INCREMENT,
		username VARCHAR(64) NULL,
		age int	NULL,
		created DATETIME default CURRENT_TIMESTAMP
	)`
	res, err := db.Exec(sqlTable)
	if err != nil {
		fmt.Println("exec error: ", err)
		return
	}
	rowNum, err := res.RowsAffected()
	if err != nil {
		fmt.Println("exec error: ", err)
		return
	}
	fmt.Println("执行数量：", rowNum)

	// 新增数据
	// var stmt, _ = db.Prepare("INSERT INTO userinfo(username, age) VALUES(?, ?)")
	// res, _ = stmt.Exec("Tom", "18")
	// id, _ := res.LastInsertId()
	// fmt.Printf("新增数据的ID：%v\n", id)

	// 批量新增数据
	// var params = []any{"Lily", "20", "Jim", "30"}
	// var stmt, _ = db.Prepare("INSERT INTO userinfo(username, age) VALUES(?, ?), (?, ?)")
	// res, _ = stmt.Exec(params...)
	// affectedNum, _ := res.RowsAffected()
	// fmt.Printf("新增数量：%v\n", affectedNum)

	// 更新数据
	// stmt, _ := db.Prepare("update userinfo set username=? where id=?")
	// res, _ = stmt.Exec("Tim", 1)
	// affectedNum, _ := res.RowsAffected()
	// fmt.Printf("更新数量：%v\n", affectedNum)

	// 删除数据
	// stmt, _ := db.Prepare("DELETE FROM userinfo WHERE id=?")
	// res, _ = stmt.Exec(1)
	// affectedNum, _ := res.RowsAffected()
	// fmt.Printf("删除数量：%v\n", affectedNum)

	// 批量删除
	// var params = []any{2, 3}
	// stmt, _ := db.Prepare("DELETE FROM userinfo WHERE id IN (?, ?)")
	// res, _ = stmt.Exec(params...)
	// affectedNum, _ := res.RowsAffected()
	// fmt.Printf("删除数量：%v\n", affectedNum)

	// 查询数据
	var users = make([]userinfo, 1)
	var rows, _ = db.Query("SELECT * FROM userinfo WHERE id=?", 4)
	var i = 0
	for rows.Next() {
		rows.Scan(&users[i].ID, &users[i].Name, &users[i].Age, &users[i].Created)
		i++
	}
	fmt.Printf("数据：%+v\n", users)

	db.Close()
}
