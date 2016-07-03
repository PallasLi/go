
package main
//这个程序是为了把log日志目录下的文件
//文件格式是*.2012-10-12.txt的
//创建2012-10-12目录
//再把文件mv到相应的目录
//Licence BSD
//Author Eagle
//Since 2012-11-08
import (
     "os"
    "os/exec"
    "io"
    "io/ioutil"
    "fmt"
    "log"
    "path/filepath"
)


//关闭文件     
func (file *File) Close() os.Error {
    if file == nil {
        return os.EINVAL
    }
    e := syscall.Close(file.fd)
    file.fd = -1 // so it can't be closed again
    if e != 0 {
        return os.Errno(e)
    }
    return nil
}
//文件读取     
func (file *File) Read(b []byte) (ret int, err os.Error) {
    if file == nil {
        return -1, os.EINVAL
    }
    r, e := syscall.Read(file.fd, b)
    if e != 0 {
        err = os.Errno(e)
    }
    return int(r), err
}
//写文件     
func (file *File) Write(b []byte) (ret int, err os.Error) {
    if file == nil {
        return -1, os.EINVAL
    }
    r, e := syscall.Write(file.fd, b)
    if e != 0 {
        err = os.Errno(e)
    }
    return int(r), err
}
//获取文件名     
func (file *File) String() string {
    return file.name
}

//遍历目录
filepath.Walk("/home/leo", 
        func(path string,f os.FileInfo, err error) error {
            if (f == nil) {
                return err
            }
            if f.IsDir() {
                return nil
            }
            println(path)
            return nil
        })   


//语言获取文件sha1值
file, err := os.Open("./file/Canon.mp3")
    if err != nil {
        return
    }
    defer file.Close()
    h := sha1.New()
    _, erro := io.Copy(h, file)
    if erro != nil {
        return
    }
    fmt.Printf("%x\n", h.Sum(nil))



func checkError(err error) {
    if err != nil {
        log.Fatalf("Error: %s", err)
    }
}
//这个是windows使用的
func moveFiles(path string){
    moveFile(path,`*.`+path+`.txt`)
}
//一个一个文件move
func moveFile(path string,filename string){
    cmd := exec.Command("mv" ,filename,path)

    // Create stdout, stderr streams of type io.Reader
    stdout, err := cmd.StdoutPipe()
    checkError(err)
    stderr, err := cmd.StderrPipe()
    checkError(err)

    // Start command
    err = cmd.Start()
    checkError(err)

    // Don't let main() exit before our command has finished running
    defer cmd.Wait()  // Doesn't block

    // Non-blockingly echo command output to terminal
    go io.Copy(os.Stdout, stdout)
    go io.Copy(os.Stderr, stderr)
    // I love Go's trivial concurrency :-D
    //fmt.Printf("Do other stuff here! No need to wait.\n\n")
}
func ListDir(dir string) ([]os.FileInfo, error) {
    return ioutil.ReadDir(dir)
}
func dirRun(fi []os.FileInfo,err error)([]string){
    if err != nil {
        fmt.Println("Error", err)
    }
    var arrPath []string;
    for _, f := range fi {
        //d := "-"
        if f.IsDir() {// d = "d" 
        }else{
            var fileName=f.Name();
            var ilen=len(fileName)
            if(string(fileName[ilen-4:ilen])==`.txt`){
                 //fmt.Printf("%s %o %d %s %s\n", d, f.Mode() & 0777, f.Size(), f.ModTime().Format("1-2 15:04"), fileName)
                 var strPath=string(fileName[ilen-14:ilen-4])
                 iFind:=false
                 for _,s:= range arrPath{
                    if(strPath==s){
                        iFind=true
                        break
                    }
                }
                if(!iFind){
                    arrPath=append(arrPath,strPath)
                }
                 
            }
            
        }
    }
    return arrPath
}
func isExists( path string ) bool {
    _,err := os.Stat( path )
    if err == nil {
        return true
    }
    return os.IsExist( err )
}
func main() {
    dir := "./"
    if len(os.Args) > 1 {
        dir = os.Args[1]
    }
    fi, err := ListDir(dir)
    paths:=dirRun(fi,err)
    for _,s:= range paths{
        if(!isExists(s)){
            os.Mkdir(s,os.ModeDir)
            fmt.Println(s)
        }
        math,_:=filepath.Glob(`*.`+s+`.txt`)
        //fmt.Println(math)
        //moveFiles(s)
        for _,t:=range math{
            moveFile(s,t)
        }
        //fmt.Println(math)
        //moveFiles(s)
    }
}

