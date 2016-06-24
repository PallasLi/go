package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter,r *http.Request) {
	fmt.Fprint(w, "hello world")
}

func main() {
	handler:=http.HandlerFunc(hello)
	http.ListenAndServe("localhost:8888", handler)
}

