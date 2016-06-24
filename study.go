package main

import (
	"fmt"
	"os"
	"imp"
	"flag"
)

const (
	USERNAME = "lyt1987"
)

var i int64
var m map[string]int

//var c chan
var i1, s1 = 123, "hello"

func testDataType() {
	i2, s2 := 123, "hello"
	var array1 [3]int = [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 3}
	array3 := []int{1: 1, 3: 2}
	array4 := [30]int{1: 2, 3: 2}
	var array5 []int = []int{1: 2, 3: 3}
	fmt.Printf("ddafd")

	fmt.Println("s2:%s", s2) //函数内定义的变量必须使用，否则编译不通过

	fmt.Println("i2:%d", i2)

	fmt.Println("s1:%s", s1)
	fmt.Println("array1[1]:%d", array1[1:2])
	fmt.Println("array2[1]:%s", array2[1:3])

	ba1 := []byte("abcdef") //字符串是不可改变的，只能转换为byte数组后进行修改
	ba1[2] = 'B'
	fmt.Println("changed:%s", string(ba1))
	fmt.Println(array1, array2, array3, array4, array5)

	//字典map
	stringMap := map[int]string{1: "A", 2: "b"}
	fmt.Println(stringMap)
	array5 = append(array5, []int{1, 2}...)
	fmt.Println(array5)
	newlineString := `adflalf\t
	afa`
	fmt.Println(newlineString)
}

func getTwoResult() (int, int) {
	return 1, 2
}

func testIf() {
	if USERNAME == "lyt1987" {
		fmt.Println("username is :%s", USERNAME)
	}
	//不使用的变量用_代替，表示抛弃
	if _, b := getTwoResult(); b == 1 {
		fmt.Println("username is :%s", USERNAME)
	}
}

func testSwitch() {
	n := 2
	switch n {
	case 2:
		fmt.Println("equal")
		fallthrough
	case 1:
		fmt.Println("not equal")
	case 0:
		fmt.Println("not equal")
	default:
		fmt.Println("other")
	}

}

func testFor() {
	//range 可处理字符串、数组、map等
	for i := range s1 {
		fmt.Println(i)
	}
}
func testOs(filePath string ) error{
	file,err:=os.Open(filePath)
	if err!=nil{
		return err
	}
	defer  func(){file.Close()}()
	fmt.Println( file.Name())
	return nil
}
var name string
func init (){
	fmt.Println("...")
	flag.StringVar(&name, "name", "value", "usage")
}
func main() {
	testDataType()
	testIf()
	testFor()
	testSwitch()
	err:=testOs("study.do")
	fmt.Println(err)
	fmt.Println(imp.Imp)
	fmt.Println(imp.PP)
	// fmt.Println(imp.notimp)
	flag.Parse()
	fmt.Println(name)
}
