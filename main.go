package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func main() {
	http.HandleFunc("/", hello)

	fmt.Println("开始监听8000端口.........")
	http.ListenAndServe(":8000", nil)
}
