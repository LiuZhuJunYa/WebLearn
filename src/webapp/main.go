package main

import (
	"fmt"
	"net/http"
)

//创建处理器函数，不需要实现Sever方法，利用HandleFunc方法可以自动生成处理器
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)

	//创建路由
	http.ListenAndServe(":8080", nil)
}
