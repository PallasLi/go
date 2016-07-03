package main

import (
	"encoding/base64"
	"fmt"
)

const (
	base64Table = "123QRSTUabcdVWXYZHijKLAWDCABDstEFGuvwxyzGHIJklmnopqr234560178912"
)

var coder = base64.NewEncoding(base64Table)

func base64Encode(src []byte) []byte {
	return []byte(coder.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
	return coder.DecodeString(string(src))
}

func testBase64(s string){
	// base64加密 
	debyte := base64Encode([]byte(s))
	fmt.Printf( "%v 加密后 %v\n",s,string(debyte) )

	// base64解密
	enbyte, err := base64Decode(debyte)
	if err != nil {
		fmt.Println(err.Error())
	}
	if s != string(enbyte) {
		fmt.Printf("v% is not equal to enbyte\n",s)
	}
	fmt.Printf( "%v 解密后 %v\n",string(debyte),string(enbyte) ) 
}

func main() {
	testBase64("hello world2")
}
