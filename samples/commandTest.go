package main

import (
	"log"
	"os"
	"os/exec"
	"time"
	"bytes"
	"fmt"
	"runtime"
	"strings"
)
import "syscall"
import "unsafe"

func testShell(){

	cmd := exec.Command("ls", "-l")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out.String())

	out2, _ := exec.Command("ls", "-la").CombinedOutput()
	fmt.Println(string(out2))

	// start
	cmd2 := exec.Command("sleep", "5")
	if err := cmd2.Start(); err != nil {
		panic(err)
	}

	// wait or timeout
	donec := make(chan error, 1)
	go func() {
		donec <- cmd2.Wait()
	}()
	select {
	case <-time.After(3 * time.Second):
		cmd2.Process.Kill()
		fmt.Println("timeout")
	case <-donec:
		fmt.Println("done")
	}
}
func testCmd() {


	var hand uintptr = uintptr(0)
	var operator uintptr = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("open")))
	var fpath uintptr = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("C:\\WINDOWS\\NOTEPAD.EXE")))
	var param uintptr = uintptr(0)
	var dirpath uintptr = uintptr(0)
	var ncmd uintptr = uintptr(1)
	shell32 := syscall.NewLazyDLL("shell32.dll")
	ShellExecuteW := shell32.NewProc("ShellExecuteW")
	_, _, _ = ShellExecuteW.Call(hand, operator, fpath, param, dirpath, ncmd)

	lf, err := os.OpenFile("angel.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)
	if err != nil {
		os.Exit(1)
	}
	defer lf.Close()
	// 日志
	l := log.New(lf, "", os.O_APPEND)
	for {
		cmd := exec.Command("C:\\WINDOWS\\NOTEPAD.EXE", "L:\\go-master\\samples\\text.txt")
		err := cmd.Start()
		if err != nil {
			l.Printf("%s 启动命令失败", time.Now().Format("2006-01-02 15:04:05"), err)

			time.Sleep(time.Second * 5)
			continue
		}
		l.Printf("%s 进程启动", time.Now().Format("2006-01-02 15:04:05"), err)
		err = cmd.Wait()
		l.Printf("%s 进程退出", time.Now().Format("2006-01-02 15:04:05"), err)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	osname:=runtime.GOOS;
	fmt.Printf("%v",osname)
    if strings.EqualFold("maxos", osname) { 
		testShell() 
	}else if strings.EqualFold("windows", osname) { 
		testCmd() 
	}
}