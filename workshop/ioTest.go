package main

import (
	"errors"
	"fmt"
	//"io"
)

func main() {
	var EOF = errors.New("EOF")
	fmt.Println(EOF)
		/**


		异常
		var EOF = errors.New("EOF")
		EOF当无法得到更多输入时，Read方法返回EOF。当函数一切正常的到达输入的结束时，就应返回EOF。如果在一个结构化数据流中EOF在不期望的位置出现了，则应返回错误ErrUnexpectedEOF或者其它给出更多细节的错误。

		var ErrClosedPipe = errors.New("io: read/write on closed pipe")
		当从一个已关闭的Pipe读取或者写入时，会返回ErrClosedPipe。

		var ErrNoProgress = errors.New("multiple Read calls return no data or error")
		某些使用io.Reader接口的客户端如果多次调用Read都不返回数据也不返回错误时，就会返回本错误，一般来说是io.Reader的实现有问题的标志。

		var ErrShortBuffer = errors.New("short buffer")
		ErrShortBuffer表示读取操作需要大缓冲，但提供的缓冲不够大。

		var ErrShortWrite = errors.New("short write")
		ErrShortWrite表示写入操作写入的数据比提供的少，却没有显式的返回错误。

		var ErrUnexpectedEOF = errors.New("unexpected EOF")
		ErrUnexpectedEOF表示在读取一个固定尺寸的块或者数据结构时，在读取未完全时遇到了EOF。


	  type Reader interface { Read(p []byte) (n int, err error)
	  }
	  Reader接口用于包装基本的读取方法。
	  Read方法读取len(p)字节数据写入p。它返回写入的字节数和遇到的任何错误。即使Read方法返回值n < len(p)，本方法在被调用时仍可能使用p的全部长度作为暂存空间。如果有部分可用数据，但不够len(p)字节，Read按惯例会返回可以读取到的数据，而不是等待更多数据。
	  当Read在读取n > 0个字节后遭遇错误或者到达文件结尾时，会返回读取的字节数。它可能会在该次调用返回一个非nil的错误，或者在下一次调用时返回0和该错误。一个常见的例子，Reader接口会在输入流的结尾返回非0的字节数，返回值err == EOF或err == nil。但不管怎样，下一次Read调用必然返回(0, EOF)。调用者应该总是先处理读取的n > 0字节再处理错误值。这么做可以正确的处理发生在读取部分数据后的I/O错误，也能正确处理EOF事件。
	  如果Read的某个实现返回0字节数和nil错误值，表示被阻碍；调用者应该将这种情况视为未进行操作。

	  type Writer interface { Write(p []byte) (n int, err error)
	  }
	  Writer接口用于包装基本的写入方法。
	  Write方法len(p) 字节数据从p写入底层的数据流。它会返回写入的字节数(0 <= n <= len(p))和遇到的任何导致写入提取结束的错误。Write必须返回非nil的错误，如果它返回的 n < len(p)。Write不能修改切片p中的数据，即使临时修改也不行。

	  type Closer interface { Close() error }
	  Closer接口用于包装基本的关闭方法。
	  在第一次调用之后再次被调用时，Close方法的的行为是未定义的。某些实现可能会说明他们自己的行为。

	  type Seeker interface { Seek(offset int64, whence int) (int64, error)
	  }
	  Seeker接口用于包装基本的移位方法。
	  Seek方法设定下一次读写的位置：偏移量为offset，校准点由whence确定：0表示相对于文件起始；1表示相对于当前位置；2表示相对于文件结尾。Seek方法返回新的位置以及可能遇到的错误。
	  移动到一个绝对偏移量为负数的位置会导致错误。移动到任何偏移量为正数的位置都是合法的，但其下一次I/O操作的具体行为则要看底层的实现。

	  type ReadCloser interface { Reader Closer }
	  ReadCloser接口聚合了基本的读取和关闭操作。

	  type ReadSeeker interface { Reader Seeker }
	  ReadSeeker接口聚合了基本的读取和移位操作。

	  type WriteCloser interface { Writer Closer }
	  WriteCloser接口聚合了基本的写入和关闭操作。

	  type WriteSeeker interface { Writer Seeker }
	  WriteSeeker接口聚合了基本的写入和移位操作。

	  type ReadWriter interface { Reader Writer }
	  ReadWriter接口聚合了基本的读写操作。

	  type ReadWriteCloser interface { Reader Writer Closer }
	  ReadWriteCloser接口聚合了基本的读写和关闭操作。

	  type ReadWriteSeeker interface { Reader Writer Seeker }
	  ReadWriteSeeker接口聚合了基本的读写和移位操作。

	  type ReaderAt interface { ReadAt(p []byte, off int64) (n int, err error)
	  }
	  ReaderAt接口包装了基本的ReadAt方法。
	  ReadAt从底层输入流的偏移量off位置读取len(p)字节数据写入p， 它返回读取的字节数(0 <= n <= len(p))和遇到的任何错误。当ReadAt方法返回值n < len(p)时，它会返回一个非nil的错误来说明为啥没有读取更多的字节。在这方面，ReadAt是比Read要严格的。即使ReadAt方法返回值 n < len(p)，它在被调用时仍可能使用p的全部长度作为暂存空间。如果有部分可用数据，但不够len(p)字节，ReadAt会阻塞直到获取len(p)个字节数据或者遇到错误。在这方面，ReadAt和Read是不同的。如果ReadAt返回时到达输入流的结尾，而返回值n == len(p)，其返回值err既可以是EOF也可以是nil。
	  如果ReadAt是从某个有偏移量的底层输入流（的Reader包装）读取，ReadAt方法既不应影响底层的偏移量，也不应被底层的偏移量影响。
	  ReadAt方法的调用者可以对同一输入流执行并行的ReadAt调用。

	  type WriterAt interface { WriteAt(p []byte, off int64) (n int, err error)
	  }
	  WriterAt接口包装了基本的WriteAt方法。
	  WriteAt将p全部len(p)字节数据写入底层数据流的偏移量off位置。它返回写入的字节数(0 <= n <= len(p))和遇到的任何导致写入提取中止的错误。当其返回值n < len(p)时，WriteAt必须放哪会一个非nil的错误。
	  如果WriteAt写入的对象是某个有偏移量的底层输出流（的Writer包装），WriteAt方法既不应影响底层的偏移量，也不应被底层的偏移量影响。
	  ReadAt方法的调用者可以对同一输入流执行并行的WriteAt调用。（前提是写入范围不重叠）

	  type ByteReader interface { ReadByte() (c byte, err error)
	  }
	  ByteReader是基本的ReadByte方法的包装。
	  ReadByte读取输入中的单个字节并返回。如果没有字节可读取，会返回错误。

	  type ByteScanner interface { ByteReader UnreadByte() error }
	  ByteScanner接口在基本的ReadByte方法之外还添加了UnreadByte方法。
	  UnreadByte方法让下一次调用ReadByte时返回之前调用ReadByte时返回的同一个字节。连续调用两次UnreadByte方法而中间没有调用ReadByte时，可能会导致错误。

	  type RuneReader interface { ReadRune() (r rune, size int, err error)
	  }
	  RuneReader是基本的ReadRune方法的包装。
	  ReadRune读取单个utf-8编码的字符，返回该字符和它的字节长度。如果没有有效的字符，会返回错误。

	  type RuneScanner interface { RuneReader UnreadRune() error }
	  RuneScanner接口在基本的ReadRune方法之外还添加了UnreadRune方法。
	  UnreadRune方法让下一次调用ReadRune时返回之前调用ReadRune时返回的同一个utf-8字符。连续调用两次UnreadRune方法而中间没有调用ReadRune时，可能会导致错误。

	  type ByteWriter interface { WriteByte(c byte) error }
	  ByteWriter是基本的WriteByte方法的包装。

	  type ReaderFrom interface { ReadFrom(r Reader) (n int64, err error)
	  }
	  ReaderFrom接口包装了基本的ReadFrom方法。
	  ReadFrom方法从r读取数据直到EOF或者遇到错误。返回值n是读取的字节数，执行时遇到的错误（EOF除外）也会被返回。

	  type WriterTo interface { WriteTo(w Writer) (n int64, err error)
	  }
	  WriterTo接口包装了基本的WriteTo方法。
	  WriteTo方法将数据写入w直到没有数据可以写入或者遇到错误。返回值n是写入的字节数，执行时遇到的任何错误也会被返回。

	  type LimitedReader struct { R Reader // 底层Reader接口 N int64 // 剩余可读取字节数 }
	  LimitedReader从R中读取数据，但限制可以读取的数据的量为最多N字节，每次调用Read方法都会更新N以标记剩余可以读取的字节数。

	  func LimitReader(r Reader, n int64) Reader
	  返回一个Reader，它从r中读取n个字节后以EOF停止。返回值接口的底层为*LimitedReader类型。

	  func (l *LimitedReader) Read(p []byte) (n int, err error)

	  type SectionReader struct { // 内含隐藏或非导出字段 }
	  SectionReader实现了对底层满足ReadAt接口的输入流某个片段的Read、ReadAt、Seek方法。

	  func NewSectionReader(r ReaderAt, off int64, n int64) *SectionReader
	  返回一个从r中的偏移量off处为起始，读取n个字节后以EOF停止的SectionReader。

	  func (s *SectionReader) Size() int64
	  Size返回该片段的字节数。

	  func (s *SectionReader) Read(p []byte) (n int, err error)

	  func (s *SectionReader) ReadAt(p []byte, off int64) (n int, err error)

	  func (s *SectionReader) Seek(offset int64, whence int) (int64, error)

	  type PipeReader struct { // 内含隐藏或非导出字段 }
	  PipeReader是一个管道的读取端。

	  func Pipe() (*PipeReader, *PipeWriter)
	  Pipe创建一个同步的内存中的管道。它可以用于连接期望io.Reader的代码和期望io.Writer的代码。一端的读取对应另一端的写入，直接在两端拷贝数据，没有内部缓冲。可以安全的并行调用Read和Write或者Read/Write与Close方法。Close方法会在最后一次阻塞中的I/O操作结束后完成。并行调用Read或并行调用Write也是安全的：每一个独立的调用会依次进行。

	  func (r *PipeReader) Read(data []byte) (n int, err error)
	  Read实现了标准Reader接口：它从管道中读取数据，会阻塞直到写入端开始写入或写入端被关闭。

	  func (r *PipeReader) Close() error
	  Close关闭读取器；关闭后如果对管道的写入端进行写入操作，就会返回(0, ErrClosedPip)。

	  func (r *PipeReader) CloseWithError(err error) error
	  CloseWithError类似Close方法，但将调用Write时返回的错误改为err。

	  type PipeWriter struct { // 内含隐藏或非导出字段 }
	  PipeWriter是一个管道的写入端。

	  func (w *PipeWriter) Write(data []byte) (n int, err error)
	  Write实现了标准Writer接口：它将数据写入到管道中，会阻塞直到读取器读完所有的数据或读取端被关闭。

	  func (w *PipeWriter) Close() error
	  Close关闭写入器；关闭后如果对管道的读取端进行读取操作，就会返回(0, EOF)。

	  func (w *PipeWriter) CloseWithError(err error) error
	  CloseWithError类似Close方法，但将调用Read时返回的错误改为err。

	  func TeeReader(r Reader, w Writer) Reader
	  TeeReader返回一个将其从r读取的数据写入w的Reader接口。所有通过该接口对r的读取都会执行对应的对w的写入。没有内部的缓冲：写入必须在读取完成前完成。写入时遇到的任何错误都会作为读取错误返回。

	  func MultiReader(readers ...Reader) Reader
	  MultiReader返回一个将提供的Reader在逻辑上串联起来的Reader接口。他们依次被读取。当所有的输入流都读取完毕，Read才会返回EOF。如果readers中任一个返回了非nil非EOF的错误，Read方法会返回该错误。

	  func MultiWriter(writers ...Writer) Writer
	  MultiWriter创建一个Writer接口，会将提供给其的数据写入所有创建时提供的Writer接口。

	  func Copy(dst Writer, src Reader) (written int64, err error)
	  将src的数据拷贝到dst，直到在src上到达EOF或发生错误。返回拷贝的字节数和遇到的第一个错误。
	  对成功的调用，返回值err为nil而非EOF，因为Copy定义为从src读取直到EOF，它不会将读取到EOF视为应报告的错误。如果src实现了WriterTo接口，本函数会调用src.WriteTo(dst)进行拷贝；否则如果dst实现了ReaderFrom接口，本函数会调用dst.ReadFrom(src)进行拷贝。

	  func CopyN(dst Writer, src Reader, n int64) (written int64, err error)
	  从src拷贝n个字节数据到dst，直到在src上到达EOF或发生错误。返回复制的字节数和遇到的第一个错误。
	  只有err为nil时，written才会等于n。如果dst实现了ReaderFrom接口，本函数很调用它实现拷贝。

	  func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)
	  ReadAtLeast从r至少读取min字节数据填充进buf。函数返回写入的字节数和错误（如果没有读取足够的字节）。只有没有读取到字节时才可能返回EOF；如果读取了有但不够的字节时遇到了EOF，函数会返回ErrUnexpectedEOF。 如果min比buf的长度还大，函数会返回ErrShortBuffer。只有返回值err为nil时，返回值n才会不小于min。

	  func ReadFull(r Reader, buf []byte) (n int, err error)
	  ReadFull从r精确地读取len(buf)字节数据填充进buf。函数返回写入的字节数和错误（如果没有读取足够的字节）。只有没有读取到字节时才可能返回EOF；如果读取了有但不够的字节时遇到了EOF，函数会返回ErrUnexpectedEOF。 只有返回值err为nil时，返回值n才会等于len(buf)。

	  func WriteString(w Writer, s string) (n int, err error)
	  WriteString函数将字符串s的内容写入w中。如果w已经实现了WriteString方法，函数会直接调用该方法。
	  **/
}
