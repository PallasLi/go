package main

import(
	"fmt"
    "os"
	"bytes" 
	"strconv"
)

func testStrconv(){
// ParseBool 将字符串转换为布尔值
// 它接受真值：1, t, T, TRUE, true, True
// 它接受假值：0, f, F, FALSE, false, False.
// 其它任何值都返回一个错误
//func ParseBool(str string) (value bool, err error)

 
	fmt.Println(strconv.ParseBool("1"))    // true
	fmt.Println(strconv.ParseBool("t"))    // true
	fmt.Println(strconv.ParseBool("T"))    // true
	fmt.Println(strconv.ParseBool("true")) // true
	fmt.Println(strconv.ParseBool("True")) // true
	fmt.Println(strconv.ParseBool("TRUE")) // true
	fmt.Println(strconv.ParseBool("TRue"))
	// false strconv.ParseBool: parsing "TRue": invalid syntax
	fmt.Println(strconv.ParseBool("0"))     // false
	fmt.Println(strconv.ParseBool("f"))     // false
	fmt.Println(strconv.ParseBool("F"))     // false
	fmt.Println(strconv.ParseBool("false")) // false
	fmt.Println(strconv.ParseBool("False")) // false
	fmt.Println(strconv.ParseBool("FALSE")) // false
	fmt.Println(strconv.ParseBool("FALse"))
	// false strconv.ParseBool: parsing "FAlse": invalid syntax
 

// FormatBool 将布尔值转换为字符串 "true" 或 "false"
//func FormatBool(b bool) string
 
	fmt.Println(strconv.FormatBool(0 < 1)) // true
	fmt.Println(strconv.FormatBool(0 > 1)) // false
 
// AppendBool 将布尔值 b 转换为字符串 "true" 或 "false"
// 然后将结果追加到 dst 的尾部，返回追加后的 []byte
//func AppendBool(dst []byte, b bool) []byte

 
	rst := make([]byte, 0)
	rst = strconv.AppendBool(rst, 0 < 1)
	fmt.Printf("%s\n", rst) // true
	rst = strconv.AppendBool(rst, 0 > 1)
	fmt.Printf("%s\n", rst) // truefalse
 
 
// ParseFloat 将字符串转换为浮点数
// s：要转换的字符串
// bitSize：指定浮点类型（32:float32、64:float64）
// 如果 s 是合法的格式，而且接近一个浮点值，
// 则返回浮点数的四舍五入值（依据 IEEE754 的四舍五入标准）
// 如果 s 不是合法的格式，则返回“语法错误”
// 如果转换结果超出 bitSize 范围，则返回“超出范围”
//func ParseFloat(s string, bitSize int) (f float64, err error)
 
	s := "0.12345678901234567890"
	f, err := strconv.ParseFloat(s, 32)
	fmt.Println(f, err)          // 0.12345679104328156
	fmt.Println(float32(f), err) // 0.12345679
	f, err = strconv.ParseFloat(s, 64)
	fmt.Println(f, err) // 0.12345678901234568
 
/**
// ErrRange 表示值超出范围
var ErrRange = errors.New("value out of range")

// ErrSyntax 表示语法不正确
var ErrSyntax = errors.New("invalid syntax")

// NumError 记录转换失败
type NumError struct {
Func string // 失败的函数名(ParseBool, ParseInt, ParseUint, ParseFloat)
Num string // 输入的值
Err error // 失败的原因(ErrRange, ErrSyntax)
}

// int 或 uint 类型的长度(32 或 64)
const IntSize = intSize 
const intSize = 32 << uint(^uint(0)>>63)

// 实现 Error 接口，输出错误信息
func (e *NumError) Error() string
 **/

// ParseInt 将字符串转换为 int 类型
// s：要转换的字符串
// base：进位制（2 进制到 36 进制）
// bitSize：指定整数类型（0:int、8:int8、16:int16、32:int32、64:int64）
// 返回转换后的结果和转换时遇到的错误
// 如果 base 为 0，则根据字符串的前缀判断进位制（0x:16，0:8，其它:10）
//func ParseInt(s string, base int, bitSize int) (i int64, err error)
 
	fmt.Println(strconv.ParseInt("123", 10, 8))
	// 123
	fmt.Println(strconv.ParseInt("12345", 10, 8))
	// 127 strconv.ParseInt: parsing "12345": value out of range
	fmt.Println(strconv.ParseInt("2147483647", 10, 0))
	// 2147483647
	fmt.Println(strconv.ParseInt("0xFF", 16, 0))
	// 0 strconv.ParseInt: parsing "0xFF": invalid syntax
	fmt.Println(strconv.ParseInt("FF", 16, 0))
	// 255
	fmt.Println(strconv.ParseInt("0xFF", 0, 0))
	// 255
 

// ParseUint 功能同 ParseInt 一样，只不过返回 uint 类型整数
// func ParseUint(s string, base int, bitSize int) (n uint64, err error)
 
	fmt.Println(strconv.ParseUint("FF", 16, 8))
	// 255
 
// Atoi 相当于 ParseInt(s, 10, 0)
// 通常使用这个函数，而不使用 ParseInt
//func Atoi(s string) (i int, err error)
 
	fmt.Println(strconv.Atoi("2147483647"))
	// 2147483647
	fmt.Println(strconv.Atoi("2147483648"))
	// 2147483647 strconv.ParseInt: parsing "2147483648": value out of range
 
 
// FormatFloat 将浮点数 f 转换为字符串值
// f：要转换的浮点数
// fmt：格式标记（b、e、E、f、g、G）
// prec：精度（数字部分的长度，不包括指数部分）
// bitSize：指定浮点类型（32:float32、64:float64）
//
// 格式标记：
// 'b' (-ddddp±ddd，二进制指数)
// 'e' (-d.dddde±dd，十进制指数)
// 'E' (-d.ddddE±dd，十进制指数)
// 'f' (-ddd.dddd，没有指数)
// 'g' ('e':大指数，'f':其它情况)
// 'G' ('E':大指数，'f':其它情况)
//
// 如果格式标记为 'e'，'E'和'f'，则 prec 表示小数点后的数字位数
// 如果格式标记为 'g'，'G'，则 prec 表示总的数字位数（整数部分+小数部分）
//func FormatFloat(f float64, fmt byte, prec, bitSize int) string
 
	f  = 100.12345678901234567890123456789
	fmt.Println(strconv.FormatFloat(f, 'b', 5, 32))
	// 13123382p-17
	fmt.Println(strconv.FormatFloat(f, 'e', 5, 32))
	// 1.00123e+02
	fmt.Println(strconv.FormatFloat(f, 'E', 5, 32))
	// 1.00123E+02
	fmt.Println(strconv.FormatFloat(f, 'f', 5, 32))
	// 100.12346
	fmt.Println(strconv.FormatFloat(f, 'g', 5, 32))
	// 100.12
	fmt.Println(strconv.FormatFloat(f, 'G', 5, 32))
	// 100.12
	fmt.Println(strconv.FormatFloat(f, 'b', 30, 32))
	// 13123382p-17
	fmt.Println(strconv.FormatFloat(f, 'e', 30, 32))
	// 1.001234588623046875000000000000e+02
	fmt.Println(strconv.FormatFloat(f, 'E', 30, 32))
	// 1.001234588623046875000000000000E+02
	fmt.Println(strconv.FormatFloat(f, 'f', 30, 32))
	// 100.123458862304687500000000000000
	fmt.Println(strconv.FormatFloat(f, 'g', 30, 32))
	// 100.1234588623046875
	fmt.Println(strconv.FormatFloat(f, 'G', 30, 32))
	// 100.1234588623046875
 
// AppendFloat 将浮点数 f 转换为字符串值，并将转换结果追加到 dst 的尾部
// 返回追加后的 []byte
//func AppendFloat(dst []byte, f float64, fmt byte, prec int, bitSize int) []byte
 
	f  = 100.12345678901234567890123456789
	b := make([]byte, 0)
	b = strconv.AppendFloat(b, f, 'f', 5, 32)
	b = append(b, " "...)
	b = strconv.AppendFloat(b, f, 'e', 5, 32)
	fmt.Printf("%s", b) // 100.12346 1.00123e+0
 
// FormatUint 将 int 型整数 i 转换为字符串形式
// base：进位制（2 进制到 36 进制）
// 大于 10 进制的数，返回值使用小写字母 'a' 到 'z'
//func FormatInt(i int64, base int) string
 
	i := int64(-2048)
	fmt.Println(strconv.FormatInt(i, 2))  // -100000000000
	fmt.Println(strconv.FormatInt(i, 8))  // -4000
	fmt.Println(strconv.FormatInt(i, 10)) // -2048
	fmt.Println(strconv.FormatInt(i, 16)) // -800
	fmt.Println(strconv.FormatInt(i, 36)) // -1kw
 
// FormatUint 将 uint 型整数 i 转换为字符串形式
// base：进位制（2 进制到 36 进制）
// 大于 10 进制的数，返回值使用小写字母 'a' 到 'z'
//func FormatUint(i uint64, base int) string
 
	i2  := uint64(2048)
	fmt.Println(strconv.FormatUint(i2, 2))  // 100000000000
	fmt.Println(strconv.FormatUint(i2, 8))  // 4000
	fmt.Println(strconv.FormatUint(i2, 10)) // 2048
	fmt.Println(strconv.FormatUint(i2, 16)) // 800
	fmt.Println(strconv.FormatUint(i2, 36)) // 1kw
 
// Itoa 相当于 FormatInt(i, 10)
//func Itoa(i int) string

 
	fmt.Println(strconv.Itoa(-2048)) // -2048
	fmt.Println(strconv.Itoa(2048))  // 2048
 
// AppendInt 将 int 型整数 i 转换为字符串形式，并追加到 dst 的尾部
// i：要转换的字符串
// base：进位制
// 返回追加后的 []byte
//func AppendInt(dst []byte, i int64, base int) []byte
 
	b  = make([]byte, 0)
	b = strconv.AppendInt(b, -2048, 16)
	fmt.Printf("%s", b) // -800
 
// AppendUint 将 uint 型整数 i 转换为字符串形式，并追加到 dst 的尾部
// i：要转换的字符串
// base：进位制
// 返回追加后的 []byte
//func AppendUint(dst []byte, i uint64, base int) []byte

 
    b  = make([]byte, 0)
    b = strconv.AppendUint(b, 2048, 16)
    fmt.Printf("%s", b) // 800
 
 
// Quote 将字符串 s 转换为“双引号”引起来的字符串
// 其中的特殊字符将被转换为“转义字符”
// “不可显示的字符”将被转换为“转义字符”
//func Quote(s string) string
 
	fmt.Println(strconv.Quote(`C:\Windows`))
	// "C:\\Windows"
 

//注：此处是反引号（键盘上１左侧那个按键），而不是单引号
 
// AppendQuote 将字符串 s 转换为“双引号”引起来的字符串，
// 并将结果追加到 dst 的尾部，返回追加后的 []byte
// 其中的特殊字符将被转换为“转义字符”
//func AppendQuote(dst []byte, s string) []byte
 
	s  = `C:\Windows`
	b  = make([]byte, 0)
	b = strconv.AppendQuote(b, s)
	fmt.Printf("%s", b) // "C:\\Windows"
 
// QuoteToASCII 将字符串 s 转换为“双引号”引起来的 ASCII 字符串
// “非 ASCII 字符”和“特殊字符”将被转换为“转义字符”
//func QuoteToASCII(s string) string

 
	fmt.Println(strconv.QuoteToASCII("Hello 世界！"))
	// "Hello \u4e16\u754c\uff01"
 
// AppendQuoteToASCII 将字符串 s 转换为“双引号”引起来的 ASCII 字符串，
// 并将结果追加到 dst 的尾部，返回追加后的 []byte
// “非 ASCII 字符”和“特殊字符”将被转换为“转义字符”
//func AppendQuoteToASCII(dst []byte, s string) []byte

 
	s  = "Hello 世界！"
	b  = make([]byte, 0)
	b = strconv.AppendQuoteToASCII(b, s)
	fmt.Printf("%s", b) // "Hello \u4e16\u754c\uff01"
 
// QuoteRune 将 Unicode 字符转换为“单引号”引起来的字符串
// “特殊字符”将被转换为“转义字符”
//func QuoteRune(r rune) string
 
    fmt.Println(strconv.QuoteRune('好'))
    // '好'
 

//注：此处为单引号，而不是反引号，这点要与Quote()使用去分开

 
// AppendQuoteRune 将 Unicode 字符转换为“单引号”引起来的字符串，
// 并将结果追加到 dst 的尾部，返回追加后的 []byte
// “特殊字符”将被转换为“转义字符”
//func AppendQuoteRune(dst []byte, r rune) []byte
 
	b  = make([]byte, 0)
	b = strconv.AppendQuoteRune(b, '好')
	fmt.Printf("%s", b) // '好'
 

// QuoteRuneToASCII 将 Unicode 字符转换为“单引号”引起来的 ASCII 字符串
// “非 ASCII 字符”和“特殊字符”将被转换为“转义字符”
//func QuoteRuneToASCII(r rune) string
 
	fmt.Println(strconv.QuoteRuneToASCII('好'))
	// '\u597d'
 
// AppendQuoteRune 将 Unicode 字符转换为“单引号”引起来的 ASCII 字符串，
// 并将结果追加到 dst 的尾部，返回追加后的 []byte
// “非 ASCII 字符”和“特殊字符”将被转换为“转义字符”
//func AppendQuoteRuneToASCII(dst []byte, r rune) []byte

 
	b  = make([]byte, 0)
	b = strconv.AppendQuoteRuneToASCII(b, '好')
	fmt.Printf("%s", b) // '\u597d'
 
// CanBackquote 判断字符串 s 是否可以表示为一个单行的“反引号”字符串
// 字符串中不能含有控制字符（除了 \t）和“反引号”字符，否则返回 false
//func CanBackquote(s string) bool
 
	b2  := strconv.CanBackquote("C:\\Windows\n")
	fmt.Println(b2) // false
	b2 = strconv.CanBackquote("C:\\Windows\r")
	fmt.Println(b2) // false
	b2 = strconv.CanBackquote("C:\\Windows\f")
	fmt.Println(b2) // false
	b2 = strconv.CanBackquote("C:\\Windows\t")
	fmt.Println(b2) // true
	b2 = strconv.CanBackquote("C:\\`Windows`")
	fmt.Println(b2) // false
 

// UnquoteChar 将 s 中的第一个字符“取消转义”并解码
//
// s：转义后的字符串
// quote：字符串使用的“引号符”（用于对引号符“取消转义”）
//
// value： 解码后的字符
// multibyte：value 是否为多字节字符
// tail： 字符串 s 除去 value 后的剩余部分
// error： 返回 s 中是否存在语法错误
//
// 参数 quote 为“引号符”
// 如果设置为单引号，则 s 中允许出现 \' 字符，不允许出现单独的 ' 字符
// 如果设置为双引号，则 s 中允许出现 \" 字符，不允许出现单独的 " 字符
// 如果设置为 0，则不允许出现 \' 或 \" 字符，可以出现单独的 ' 或 " 字符
//func UnquoteChar(s string, quote byte) (value rune, multibyte bool, tail string, err error)

 
	s2 := `\"大\\家\\好！\"`
	c, mb, sr, _ := strconv.UnquoteChar(s2, '"')
	fmt.Printf("%-3c %v\n", c, mb)
	for ; len(sr) > 0; c, mb, sr, _ = strconv.UnquoteChar(sr, '"') {
		fmt.Printf("%-3c %v\n", c, mb)
	}
	// " false
	// 大 true
	// \ false
	// 家 true
	// \ false
	// 好 true
	// ！ true
 

// Unquote 将“带引号的字符串” s 转换为常规的字符串（不带引号和转义字符）
// s 可以是“单引号”、“双引号”或“反引号”引起来的字符串（包括引号本身）
// 如果 s 是单引号引起来的字符串，则返回该该字符串代表的字符
//func Unquote(s string) (t string, err error)
 
	sr2, err := strconv.Unquote(`"\"大\t家\t好！\""`)
	fmt.Println(sr2, err)
	sr2, err = strconv.Unquote(`'大家好！'`)
	fmt.Println(sr2, err)
	sr2, err = strconv.Unquote(`'好'`)
	fmt.Println(sr2, err)
	sr2, err = strconv.Unquote("`大\\t家\\t好！`")
	fmt.Println(sr2, err)
 
// IsPrint 判断 Unicode 字符 r 是否是一个可显示的字符
// 可否显示并不是你想象的那样，比如空格可以显示，而\t则不能显示
// 具体可以参考 Go 语言的源码
//func IsPrint(r rune) bool
 
	fmt.Println(strconv.IsPrint('a'))  // true
	fmt.Println(strconv.IsPrint('好'))  // true
	fmt.Println(strconv.IsPrint(' '))  // true
	fmt.Println(strconv.IsPrint('\t')) // false
	fmt.Println(strconv.IsPrint('\n')) // false
	fmt.Println(strconv.IsPrint(0))    // false
 

}



func testBuffer(){

	/**
		bytes.buffer是一个缓冲byte类型的缓冲器，这个缓冲器里存放着都是byte
		创建一个缓冲器,有三种方式 NewBuffer NewBufferString 
		 
		写入到缓冲器（缓冲器变大）
		使用Write方法，将一个byte类型的slice放到缓冲器的尾部
		Write---- func (b *Buffer) Write(p []byte) (n int, err error)
		使用WriteString方法，将一个字符串放到缓冲器的尾部
		WriteString---- func (b *Buffer) WriteString(s string) (n int, err error)
		使用WriteByte方法，将一个byte类型的数据放到缓冲器的尾部
		WriteByte---- func (b *Buffer) WriteByte(c byte) error
		使用WriteRune方法，将一个rune类型的数据放到缓冲器的尾部
		WriteRune---- func (b *Buffer) WriteRune(r rune) (n int, err error)

		从缓冲器写出（缓冲器变小）
		使用WriteTo方法，将一个缓冲器的数据写到w里，w是实现io.Writer的，比如os.File就是实现io.Writer
		WriteTo---- func (b *Buffer) WriteTo(w io.Writer) (n int64, err error)

		读出缓冲器（缓冲器变小）
		给Read方法一个容器p，读完后，p就满了，缓冲器相应的减少了，返回的n为成功读的数量
		Read----func (b *Buffer) Read(p []byte) (n int, err error)

		返回缓冲器头部的第一个byte，缓冲器头部第一个byte被拿掉
		ReadByte---- func (b *Buffer) ReadByte() (c byte, err error)

		ReadRune和ReadByte很像 返回缓冲器头部的第一个rune，缓冲器头部第一个rune被拿掉
		ReadRune---- func (b *Buffer) ReadRune() (r rune, size int, err error)
		 
		ReadBytes需要一个byte作为分隔符，读的时候从缓冲器里找第一个出现的分隔符（delim），找到后，把从缓冲器头部开始到分隔符之间的所有byte进行返回，作为byte类型的slice，返回后，缓冲器也会空掉一部分
		ReadBytes---- func (b *Buffer) ReadBytes(delim byte) (line []byte, err error)
		 
		ReadBytes需要一个byte作为分隔符，读的时候从缓冲器里找第一个出现的分隔符（delim），找到后，把从缓冲器头部开始到分隔符之间的所有byte进行返回，作为字符串，返回后，缓冲器也会空掉一部分
		ReadString---- func (b *Buffer) ReadString(delim byte) (line string, err error)
		 
		读入缓冲器（缓冲器变大）
		从一个实现io.Reader接口的r，把r里的内容读到缓冲器里，n返回读的数量
		ReadFrom---- func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error)

		从缓冲器取出（缓冲器变小）
		返回前n个byte，成为slice返回，原缓冲器变小
		Next ---- func (b *Buffer) Next(n int) []byte
	**/

	buf1:=bytes.NewBufferString("hello")
	buf2:=bytes.NewBuffer([]byte("hello"))
	buf3:=bytes.NewBuffer([]byte{'h','e','l','l','o'})
    fmt.Println(buf1.String()) 
    fmt.Println(buf2.String()) 
	fmt.Println(buf3.String()) 
	
	
    s := " world"
    s2 := []byte(" world")
    var s3 byte = '!'
    var s4 rune = '好' 
    buf := bytes.NewBufferString("hello")
    fmt.Println(buf.String())  //buf.String()方法是吧buf里的内容转成string，以便于打印
    buf.Write(s2) //将s2这个slice写到buf的尾部
    buf.WriteString(s ) //将s这个string写到buf的尾部
    buf.WriteByte(s3) //将s3这个string写到buf的尾部
    buf.WriteRune(s4) //将s4这个string写到buf的尾部 
    fmt.Println(buf.String())  //打印 hello world
   
    file , _ := os.Create("text.txt") 
    buf.WriteTo(file) //hello写到text.txt文件中了
    fmt.Fprintf(file, buf.String()) //虽然这不在讨论范围，但这句效果同上
  
    ss1:=[]byte("hello")                //申明一个slice为ss1
    buff:=bytes.NewBuffer(ss1)     //new一个缓冲器buff，里面存着hello这5个byte
    ss2:=[]byte(" world")                 //申明另一个slice为ss2
    buff.Write(ss2)                     //把ss2写入添加到buff缓冲器内
    fmt.Println(buff.String())            //使用缓冲器的String方法转成字符串，并打印："hello world"

    ss3:=make([]byte,3)               //申明一个空的slice为ss3，容量为3
    buff.Read(ss3)                       //把buff的内容读入到ss3内，因为ss3的容量为3，所以只读了3个过来
    fmt.Println(buff.String())       //buff的前3个字符被读走了，所以buff变成："lo world"
    fmt.Println(string(ss3))          //空的ss3被写入3个字符，所以为"hel"
    buff.Read(ss3)                       //把buff的内容读入到ss3内，因为ss3的容量为3，所以只读了3个过来，原来ss3的内容被覆盖了
    fmt.Println(buff.String())       //buff的前3个字符又被读走了，所以buff变成："world"
    fmt.Println(string(ss3))          //原来的ss3被从"hel"变成"lo "，因为"hel"被覆盖了
   
    fmt.Println(buf.String()) //buf.String()方法是吧buf里的内容转成string，>以便于打印
    b, _ := buf.ReadByte()    //读取第一个byte，赋值给b
    fmt.Println(buf.String()) //打印 ello，缓冲器头部第一个h被拿掉
    fmt.Println(string(b))    //打印 h
   
    fmt.Println(buf.String()) //buf.String()方法是吧buf里的内容转成string，>以便于打印
    b2, n, _ := buf.ReadRune() //读取第一个rune，赋值给b
    fmt.Println(buf.String()) //打印 hello
    fmt.Println(string(b2))    //打印中文字： 好，缓冲器头部第一个“好”被拿掉
    fmt.Println(n)            //打印3，“好”作为utf8储存占3个byte
    b2, n, _ = buf.ReadRune()  //再读取第一个rune，赋值给b
    fmt.Println(buf.String()) //打印 ello
    fmt.Println(string(b2))    //打印h，缓冲器头部第一个h被拿掉
    fmt.Println(n)            //打印 1，“h”作为utf8储存占1个byte
 
    var d byte = 'e' //分隔符为e 
    fmt.Println(buf.String()) //buf.String()方法是吧buf里的内容转成string，以便于打印
    b3, _ := buf.ReadBytes(d)  //读到分隔符，并返回给b
    fmt.Println(buf.String()) //打印 llo，缓冲器被取走一些数据
    fmt.Println(string(b3))    //打印 he，找到e了，将缓冲器从头开始，到e的内容都返回给b
   
    fmt.Println(buf.String()) //buf.String()方法是吧buf里的内容转成string，以便于打印
    b4, _ := buf.ReadString(d)  //读到分隔符，并返回给b
    fmt.Println(buf.String()) //打印 llo，缓冲器被取走一些数据
    fmt.Println(b4)    //打印 he，找到e了，将缓冲器从头开始，到e的内容都返回给b
  
    file2, _ := os.Open("test.txt")  //test.txt的内容是“world” 
    buf.ReadFrom(file2)              //将text.txt内容追加到缓冲器的尾部
    fmt.Println(buf.String())    //打印“hello world”
  
    fmt.Println(buf.String())
    b5 := buf.Next(2)   //重头开始，取2个
    fmt.Println(buf.String())  //变小了
    fmt.Println(string(b5))   //打印he
  
}


func main(){
	testBuffer() 
}
