// goTest project main.go
package main

import (
	"fmt"
	"math/big"
)

var a = 1234
var b string = "afaf"
var c bool

func test() (d *big.Int) {
	d = big.NewInt(32937843487873874)
	e := big.NewInt(32937843487873874)
	mapvar := map[string]int{"key1": 1, "key2": 2}
	println(mapvar)
	//	someArray :=[1,2,3]
	//	println(someArray)
	//	rangevari,rangervars := range a{
	//		println(rangevari,rangervars)
	//	}

	println(d.Int64)
	println(e.Int64)
	d = d.Add(d, e)
	println(d.Int64)
	f, g := 0, 1 //:= 不需要var 和类型  但是在方法外使用报错Non-declaration statement outside function body
	println(f, g)
	return
}

func main() {

	fmt.Println("hello go")
	println(a, b, c)
	test()

}
