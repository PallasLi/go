package url

import "net/url"

func main() {

	/**
	  func QueryEscape(s string) string
	  QueryEscape函数对s进行转码使之可以安全的用在URL查询里。

	  func QueryUnescape(s string) (string, error)
	  QueryUnescape函数用于将QueryEscape转码的字符串还原。它会把%AB改为字节0xAB，将'+'改为' '。如果有某个%后面未跟两个十六进制数字，本函数会返回错误。

	  type Error struct { Op string URL string Err error }
	  Error会报告一个错误，以及导致该错误发生的URL和操作。

	  func (e *Error) Error() string

	  type EscapeError string

	  func (e EscapeError) Error() string

	  type URL struct { Scheme string Opaque string // 编码后的不透明数据 User *Userinfo // 用户名和密码信息 Host string // host或host:port Path string RawQuery string // 编码后的查询字符串，没有'?' Fragment string // 引用的片段（文档位置），没有'#' }
	  URL类型代表一个解析后的URL（或者说，一个URL参照）。URL基本格式如下：

	  scheme://[userinfo@]host/path[?query][#fragment]
	  scheme后不是冒号加双斜线的URL被解释为如下格式：

	  scheme:opaque[?query][#fragment]
	  注意路径字段是以解码后的格式保存的，如/%47%6f%2f会变成/Go/。这导致我们无法确定Path字段中的斜线是来自原始URL还是解码前的%2f。除非一个客户端必须使用其他程序/函数来解析原始URL或者重构原始URL，这个区别并不重要。此时，HTTP服务端可以查询req.RequestURI，而HTTP客户端可以使用URL{Host: "example.com", Opaque: "//example.com/Go%2f"}代替{Host: "example.com", Path: "/Go/"}。

	  func Parse(rawurl string) (url *URL, err error)
	  Parse函数解析rawurl为一个URL结构体，rawurl可以是绝对地址，也可以是相对地址。

	  func ParseRequestURI(rawurl string) (url *URL, err error)
	  ParseRequestURI函数解析rawurl为一个URL结构体，本函数会假设rawurl是在一个HTTP请求里，因此会假设该参数是一个绝对URL或者绝对路径，并会假设该URL没有#fragment后缀。（网页浏览器会在去掉该后缀后才将网址发送到网页服务器）

	  func (u *URL) IsAbs() bool
	  函数在URL是绝对URL时才返回真。

	  func (u *URL) Query() Values
	  Query方法解析RawQuery字段并返回其表示的Values类型键值对。

	  func (u *URL) RequestURI() string
	  RequestURI方法返回编码好的path?query或opaque?query字符串，用在HTTP请求里。

	  func (u *URL) String() string
	  String将URL重构为一个合法URL字符串。

	  func (u *URL) Parse(ref string) (*URL, error)
	  Parse方法以u为上下文来解析一个URL，ref可以是绝对或相对URL。
	  本方法解析失败会返回nil, err；否则返回结果和ResolveReference一致。

	  func (u *URL) ResolveReference(ref *URL) *URL
	  本方法根据一个绝对URI将一个URI补全为一个绝对URI，参见RFC 3986 节 5.2。参数ref可以是绝对URI或者相对URI。ResolveReference总是返回一个新的URL实例，即使该实例和u或者ref完全一样。如果ref是绝对URI，本方法会忽略参照URI并返回ref的一个拷贝。

	  type Userinfo struct { // 内含隐藏或非导出字段 }
	  Userinfo类型是一个URL的用户名和密码细节的一个不可修改的封装。一个真实存在的Userinfo值必须保证有用户名（但根据 RFC 2396可以是空字符串）以及一个可选的密码。

	  func User(username string) *Userinfo
	  User函数返回一个用户名设置为username的不设置密码的*Userinfo。

	  func UserPassword(username, password string) *Userinfo
	  UserPassword函数返回一个用户名设置为username、密码设置为password的*Userinfo。

	  这个函数应该只用于老式的站点，因为风险很大，不建议使用，参见RFC 2396。

	  func (u *Userinfo) Username() string
	  Username方法返回用户名。

	  func (u *Userinfo) Password() (string, bool)
	  如果设置了密码返回密码和真，否则会返回假。

	  func (u *Userinfo) String() string
	  String方法返回编码后的用户信息，格式为"username[:password]"。

	  type Values map[string][]string
	  Values将建映射到值的列表。它一般用于查询的参数和表单的属性。不同于http.Header这个字典类型，Values的键是大小写敏感的。

	  func ParseQuery(query string) (m Values, err error)
	  ParseQuery函数解析一个URL编码的查询字符串，并返回可以表示该查询的Values类型的字典。本函数总是返回一个包含了所有合法查询参数的非nil字典，err用来描述解码时遇到的（如果有）第一个错误。

	  func (v Values) Get(key string) string
	  Get会获取key对应的值集的第一个值。如果没有对应key的值集会返回空字符串。获取值集请直接用map。

	  func (v Values) Set(key, value string)
	  Set方法将key对应的值集设为只有value，它会替换掉已有的值集。

	  func (v Values) Add(key, value string)
	  Add将value添加到key关联的值集里原有的值的后面。

	  func (v Values) Del(key string)
	  Del删除key关联的值集。

	  func (v Values) Encode() string
	  Encode方法将v编码为url编码格式("bar=baz&foo=quux")，编码时会以键进行排序。
	  **/
}
