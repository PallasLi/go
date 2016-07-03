package main

import ( 
	"fmt"
	"net/http" 
	"io/ioutil"
	// "io"
	"log" 
	"os"
	"bytes"
)



// 1MB
const MAX_MEMORY = 1 * 1024 * 1024

func upload(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(MAX_MEMORY); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusForbidden)
	}

	for key, value := range r.MultipartForm.Value {
		fmt.Fprintf(w, "%s:%s ", key, value)
		log.Printf("%s:%s", key, value)
	}

	for _, fileHeaders := range r.MultipartForm.File {
		for _, fileHeader := range fileHeaders {
			file, _ := fileHeader.Open()
			path := fmt.Sprintf("files/%s", fileHeader.Filename)
			buf, _ := ioutil.ReadAll(file)
			ioutil.WriteFile(path, buf, os.ModePerm)
		}
	}
}

func download(w http.ResponseWriter, r *http.Request) {
	tr := &http.Transport{}
	tr.RegisterProtocol("file", http.NewFileTransport(http.Dir(".")))
	c := &http.Client{Transport: tr}
	resp, _ := c.Get("file:///main.go")
	//io.Copy(os.Stdout, resp.Body)//在服务器控制台打印,加上此句后后边内容无法输入到页面
	fmt.Printf("%v",resp.Body)
	body, _ := ioutil.ReadAll(resp.Body)
	buf := bytes.NewBuffer(body) 
	fmt.Print(buf)
	fmt.Fprint(w,buf)
}
 
func main() {
	fmt.Println(http.Dir("."))//.表示当前目录
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/download", download)
	http.Handle("/", http.FileServer(http.Dir("static")))
	err := http.ListenAndServe(":7001", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

