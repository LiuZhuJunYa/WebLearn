package main

import (
	"fmt"
	"net/http"
)

//创建处理器函数，不需要实现Sever方法，利用HandleFunc方法可以自动生成处理器
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "通过自己创建的多路复用器处理请求......", r.URL.Path)
}

func main() {
	//自己创建一个多路复用器
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)

	//创建路由
	http.ListenAndServe(":8080", mux)
}
