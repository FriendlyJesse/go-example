package example

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func body(w http.ResponseWriter, r *http.Request) {
	// 获取请求头信息
	var header = r.Header.Get("Accept-Encoding")
	fmt.Println("请求头 Accept-Encoding：", header)

	if r.Method == "GET" {
		r.ParseForm()
		fmt.Println("获取参数的方法1：", r.Form.Get("name"))
		fmt.Println("获取参数的方法2：", r.URL.Query())
		fmt.Println("获取参数的方法3：", r.FormValue("name"))
		fmt.Fprintln(w, "This is GET")
	} else {
		// 使用 form 和 postform 方法之前必须调用 parseForm 方法
		r.ParseForm()
		fmt.Println("Postform() 获取参数：", r.Form.Get("name"))
		// 与 get 一样的方法
		fmt.Println("FormValue() 获取参数：", r.FormValue("name"))
		// PostFormValue 将 postform 功能优化
		var pfv = r.PostFormValue("name")
		fmt.Println("PostFormValue() 获取参数：", pfv)

		// 获取 post 的文件数据
		r.ParseMultipartForm(1024)
		fmt.Println("MultipartForm() 获取文件数据：", r.MultipartForm)

		// FormFile() 获取上传的文件
		file, handle, _ := r.FormFile("file")
		fmt.Println("FormFile() 获取文件数据：", file, handle)

		// 接收 post 的json 数据
		con, _ := io.ReadAll(r.Body)
		fmt.Println("json 数据：", file, string(con))
		// 响应内容
		fmt.Fprintln(w, "This is POST")
	}
}

// HTML 响应内容
func indexExample(w http.ResponseWriter, r *http.Request) {
	str := `<html><head><title>My Go</title></head><body><h1>Hello World</h1></body></html>`
	w.Write([]byte(str))
}

// WriteHeader设置响应状态码
func errorExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	str :=
		`<html>
	<head><title>My Go</title></head><body><h1>Hello World</h1></body></html>`
	w.Write([]byte(str))
}

// 在Header中设置参数Location
// 并使用WriteHeader设置302状态码，即可实现URL重定向
// 重定向的URL为参数Location的参数值
func redirectExamp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://www.baidu.com/")
	w.WriteHeader(302)
}

// 定义结构体Post，用于生成JSON数据
type Post struct {
	User    string
	Threads []string
}

// 在Header中设置参数Content-Type
// 参数值为application/json，将响应内容以JSON表示
// 使用结构体Post生成JSON数据
// 由Write方法将JSON数据作为响应内容输出
func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "Go",
		Threads: []string{"first", "second", "third"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}

func cookieExample(w http.ResponseWriter, r *http.Request) {
	// 获取HTTP请求的Cookie
	c, _ := r.Cookie("csrftoken") //获取Cookie某个属性值
	fmt.Printf("获取HTTP请求的Cookie：%v\n", c)
	// 设置响应内容的Cookie
	cookie := &http.Cookie{
		Name:   "sessionid",
		Value:  "lkjsdfklsjfklsfdsfdjslf",
		MaxAge: 3600,
		Domain: "localhost",
		Path:   "/"}
	http.SetCookie(w, cookie)
	w.Write([]byte("This is Cookie"))
}

func ExecHTTP() {
	server := http.Server{
		Addr: "127.0.0.1:8080"}
	http.HandleFunc("/", indexExample)
	http.HandleFunc("/error", errorExample)
	http.HandleFunc("/redirect", redirectExamp)
	http.HandleFunc("/json", jsonExample)
	http.HandleFunc("/cookie", cookieExample)
	server.ListenAndServe()
}
