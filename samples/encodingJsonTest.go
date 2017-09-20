package main

import (
	"encoding/json"
	"os"
	"fmt"
)

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Book struct {
	Id         int        `json:"id"`
	Name       string     `json:"name"`
	Categories []Category `json:"categories"`
}
const text1 = `
{
	"foo": "hello",
	"bar": "golang!"
}
`

func testSimpleUnMarshal() {
	var m = make(map[string]string)
	json.Unmarshal([]byte(text1), &m)
	fmt.Println(m["foo"], m["bar"])
}
const text = `
{
	"id": 1,
	"name": "go语言",
	"categories": [
		{ "id": 3, "name": "简介" },
		{ "id": 4, "name": "技术" }
	]
}
`
 

func testUnMarshal() {
	var book Book
	json.Unmarshal([]byte(text), &book)
	fmt.Printf("%v\n", book) 
}

func testMarshal(){
	book :=  Book{
		Id: 1,
		Name: "go语言",
		Categories: []Category{
			{ Id: 3, Name: "简介" },
			{ Id: 4, Name: "技术" },
		},
	}
	json.NewEncoder(os.Stdout).Encode(&book)
	
}

func main() {
	testMarshal()
	testSimpleUnMarshal()
	testUnMarshal()
}
