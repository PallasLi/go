package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"math/rand"
	"time"
)

func input(ans chan int) {
	r := bufio.NewReader(os.Stdin)
	for {

		line, _, err := r.ReadLine()
		if err != nil {
			fmt.Print(">")
			continue
		}

		str := strings.Replace(string(line), "\r", "", -1)
		n, err := strconv.Atoi(str)
		if err != nil {
			fmt.Print(">")
			continue
		}
		ans <- n
	}
}

func main() {
	/**
		c:=make(chan type,vlength)   创建管道，定义了管道中传递的值类型和值数量（缓冲数量，要先释放后输入信息，否则报错）
		<-c  从管道输出信息
		c<-  往管道输入信息
	**/
	c := make(chan int, 3) //修改2为>3就报错，修改2为<3可以正常运行
	c <- 1
	c <- 2
	c <- 3
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	
	d := make(chan string, 1)  
	d <- "string1" 
	fmt.Println(<-d) 
	
	
	ans := make(chan int)
	go input(ans)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := rnd.Intn(100)
	m := rnd.Intn(100)
	fmt.Printf("  %2d + %2d = ",   n, m)
	//select 会阻塞进程  ， 当case中某一条件满足时 ，执行那个分支， 多个条件满足时随机执行一个分支
	select {
		case nm := <-ans:
			if nm == n+m {
				fmt.Println(">> correct!") 
			} else {
				fmt.Println(">> wrong!!")
			} 
		case <-time.After(5 * time.Second):
			fmt.Println()
			fmt.Println(">> timed out")
		}
	
}
