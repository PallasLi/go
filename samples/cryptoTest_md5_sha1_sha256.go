package main

import (
 "fmt"
 "crypto/md5"
 "crypto/sha1"
 "crypto/sha256"
 "io"
)
//对字符串进行MD5哈希
func a(data string) string {
 t := md5.New();
 io.WriteString(t,data);
 return fmt.Sprintf("%x",t.Sum(nil));
}
//对字符串进行SHA1哈希
func b(data string) string {
 t := sha1.New();
 io.WriteString(t,data);
 return fmt.Sprintf("%x",t.Sum(nil));
}
//对字符串进行SHA256哈希
func c(data string) string {
 t := sha256.New();
 io.WriteString(t,data);
 return fmt.Sprintf("%x",t.Sum(nil));
}
func main(){
 var data string = "abc";
 fmt.Printf("MD5 : %s\n",a(data));
 fmt.Printf("SHA1 : %s\n",b(data));
 fmt.Printf("SHA256 : %s\n",c(data));
}