package main

import (
	"fmt" 
	"bytes" 
	"io/ioutil"
	"log"  
	"net/http"
)

func getremote(w http.ResponseWriter,r *http.Request) { 
	resp, err := http.Get("http://localhost:8888/hello1")

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("%s", resp)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("%s", body) 
	buf := bytes.NewBuffer(body)
	fmt.Fprint(w, buf)

}

func hello(w http.ResponseWriter,r *http.Request) {
	fmt.Fprint(w, "hello world")
}
func hello1(w http.ResponseWriter,r *http.Request) {
	fmt.Fprint(w, "<br>hello world<br>hello world<br>hello world<br>hello world<br>hello world")
}

func main() { 
	//handler:=http.HandlerFunc("hello)
	//http.Handle("/hello",handler)
	//http.ListenAndServe(server, handler)
	http.HandleFunc("/hello",hello)
	http.HandleFunc("/hello1",hello1)
	http.HandleFunc("/getremote",getremote)
	fmt.Println("fmt:启动web服务器：localhost:8888")
	log.Println("log:启动web服务器：localhost:8888")
	http.ListenAndServe("localhost:8888", nil) //多分枝时处理函数设为null 
}
 
 