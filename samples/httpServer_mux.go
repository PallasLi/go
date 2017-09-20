package main

import (
	"fmt"  
	"log" 
	"io" 
	"net/http"
)
 
func main() { 
	
	//测试路由方式2
    mux := http.NewServeMux()
    mux.HandleFunc("/h", func(w http.ResponseWriter, r *http.Request) {
        io.WriteString(w, "hello")
    })
    mux.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
        io.WriteString(w, "byebye")
    })
    mux.HandleFunc("/hello", sayhello)
	log.Println("fmt:启动web服务器：localhost:8889/  路由有 /h /bye /hello")
	fmt.Println("fmt:启动web服务器：localhost:8889/  路由有 /h /bye /hello")
    http.ListenAndServe(":8889", mux)
}
 

func sayhello(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "hello world")
}