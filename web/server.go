package web

import (
	"fmt"
	"log"
	"net/http"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form) //解析参数
	fmt.Println("path: ", r.URL.Path)
	fmt.Println("schema: ", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key: ", k, " value: ", v)
	}

	fmt.Fprintf(w, "Hello, welcomde to go web ^_^")
}

func Test() {
	http.HandleFunc("/", sayHelloName)       //设置访问路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
