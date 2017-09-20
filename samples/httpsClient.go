package main

import (
    "fmt"
	"crypto/tls"
    "io/ioutil"
    "net/http"
)

func main() {
	//通过设置tls.Config的InsecureSkipVerify为true，client将不再对服务端的证书进行校验。
	tr := &http.Transport{
        TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
    }
    client := &http.Client{Transport: tr}
    resp, err := client.Get("https://localhost:8889")
    //resp, err := http.Get("https://localhost:8889")  //普通方式调用https
    if err != nil {
        fmt.Println("error:", err)
        return
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
}