package main

import (
    "os"
    "fmt"
    "path/filepath"
    "crypto/md5"
    "io"
)

func hashFile(path string) error {
    file, err := os.Open(path)
    h := md5.New()

    if err != nil {
        return err
    }

    defer file.Close()

    data := make([]byte, 1024)
    for{
        n, err := file.Read(data)

        if n != 0 {
            io.WriteString(h, string(data)) 
        } else {
            break
        }

        if err != nil && err != io.EOF {
            panic(err)
        }
    }

    dir := filepath.Dir(path)
    hashName := fmt.Sprintf("%x", h.Sum(nil))

    fmt.Println(path)
    os.Rename(path, filepath.Join(dir, hashName))
    return nil 
}

func visit(path string, info os.FileInfo, err error) error {
   
    if ! info.IsDir(){
       hashFile(path)
    }
    
    
    return nil
}

func main(){
    
    err := filepath.Walk("./sp", visit)

    if err != nil {

        fmt.Println(err)
    }
}