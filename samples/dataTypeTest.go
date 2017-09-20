package main

import (
	"fmt"
)

//枚举
const (
    Sunday = "星期日"
    Monday = "星期一"
    Tuesday = "星期二"
    Wednesday = "星期三"
    Thursday = "星期四"
    Friday = "星期五"
    Saturday = "星期六"
)

const (
	USERNAME = "lyt1987"
)

var i int64
var m map[string]int

//var c chan
var i1, s1 = 123, "hello"
 
func main(){

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



	/**查看(文档)包定义
	go doc packagename
	**/
	//go doc fmt

	/**
	整数:int,int8,int16,int32(rune),int64 和 uint8(byte),uint16,uint32,uint64。
	虚数:complex64和complex128
	内建类型：error
	**/
	/**定义常量
	const name type = value
	const (
		name type = value
		...
	)
	**/
	const firstconst complex64 = 5+5i
	const (
		secondconst   = 2
		thirdconst  = 5+5i
	)

	/**定义变量
	var name type
	var (
		name type
		...
	)**/
	var firstvar int64=1
	var (
		secondvar = 2
		thirdvar = 3
	)

	/**
	使用 := 不用使用var关键字也不用声明类型,但只能在方法体中才能使用
	多值赋值 name1,name2... :=value1,value2...
	不使用的变量用_代替，表示抛弃
	**/
	forthvar,fifthvar,_:=4,5,6

	/**
	多行字符串两种方式
	1. 双引号 +  ，+号必须出现在行尾，仅为字符串的拼接 
	2. 反引号，换行格式化
	**/
	firstStr:="firstline"+
	"secondline"
	secondStr:=`firstline
	secondline
	`

	/**
	fmt.Printf:  %v 全拿格式    %d 整数    %x      %T 时间
	**/
	fmt.Printf("firstconst Value is: %v\n", firstconst)
	fmt.Printf("secondconst Value is: %v\n", secondconst)
	fmt.Printf("thirdconst Value is: %v\n", thirdconst)
	fmt.Printf("firstvar Value is: %v\n", firstvar)
	fmt.Printf("secondvar Value is: %v\n", secondvar)
	fmt.Printf("thirdvar Value is: %v\n", thirdvar)
	fmt.Printf("forthvar Value is: %v\n", forthvar)
	fmt.Printf("fifthvar Value is: %v\n", fifthvar)
	fmt.Printf("firstStr Value is: %v\n", firstStr)
	fmt.Printf("secondStr Value is: %v\n", secondStr) 
}