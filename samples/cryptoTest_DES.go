package main

import (
	"encoding/base64"
	"fmt"
)
type Block interface { 
    // BlockSize returns the cipher's block size. 
    BlockSize() int 
    // Encrypt encrypts the first block in src into dst. 
    // Dst and src may point at the same memory. 
    Encrypt(dst, src []byte) 
    // Decrypt decrypts the first block in src into dst. 
    // Dst and src may point at the same memory. 
    Decrypt(dst, src []byte) 
}
type BlockMode interface {
    // BlockSize returns the mode's block size.
    BlockSize() int

    // CryptBlocks encrypts or decrypts a number of blocks. The length of
    // src must be a multiple of the block size. Dst and src may point to
    // the same memory.
    CryptBlocks(dst, src []byte)
}
func NewCBCDecrypter(b Block, iv []byte) BlockMode
func NewCBCEncrypter(b Block, iv []byte) BlockMode

type Stream interface {
    // XORKeyStream XORs each byte in the given slice with a byte from the
    // cipher's key stream. Dst and src may point to the same memory.
    XORKeyStream(dst, src []byte)
}

func DesEncrypt(origData, key []byte) ([]byte, error) {
     block, err := des.NewCipher(key)
     if err != nil {
          return nil, err
     }
     origData = PKCS5Padding(origData, block.BlockSize())
     // origData = ZeroPadding(origData, block.BlockSize())
     blockMode := cipher.NewCBCEncrypter(block, key)
     crypted := make([]byte, len(origData))
      // 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
     // crypted := origData
     blockMode.CryptBlocks(crypted, origData)
     return crypted, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
     padding := blockSize - len(ciphertext)%blockSize
     padtext := bytes.Repeat([]byte{byte(padding)}, padding)
     return append(ciphertext, padtext...)
}

func DesDecrypt(crypted, key []byte) ([]byte, error) {
     block, err := des.NewCipher(key)
     if err != nil {
          return nil, err
     }
     blockMode := cipher.NewCBCDecrypter(block, key)
     origData := make([]byte, len(crypted))
     // origData := crypted
     blockMode.CryptBlocks(origData, crypted)
     origData = PKCS5UnPadding(origData)
     // origData = ZeroUnPadding(origData)
     return origData, nil
}

// 3DES加密
func TripleDesEncrypt(origData, key []byte) ([]byte, error) {
     block, err := des.NewTripleDESCipher(key)
     if err != nil {
          return nil, err
     }
     origData = PKCS5Padding(origData, block.BlockSize())
     // origData = ZeroPadding(origData, block.BlockSize())
     blockMode := cipher.NewCBCEncrypter(block, key[:8])
     crypted := make([]byte, len(origData))
     blockMode.CryptBlocks(crypted, origData)
     return crypted, nil
}

func NewCBCDecrypter(b Block, iv []byte) BlockMode {
     if len(iv) != b.BlockSize() {
          panic("cipher.NewCBCDecrypter: IV length must equal block size")
     }
     return (*cbcDecrypter)(newCBC(b, iv))
}

// 3DES解密
func TripleDesDecrypt(crypted, key []byte) ([]byte, error) {
     block, err := des.NewTripleDESCipher(key)
     if err != nil {
          return nil, err
     }
     blockMode := cipher.NewCBCDecrypter(block, key[:8])
     origData := make([]byte, len(crypted))
     // origData := crypted
     blockMode.CryptBlocks(origData, crypted)
     origData = PKCS5UnPadding(origData)
     // origData = ZeroUnPadding(origData)
     return origData, nil
}


func main() { 
}
