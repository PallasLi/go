package goroutine

import(
	"testing"
	"runtime"
)

const(
	LOOP_COUNT=1000000000
)

//两个协程循环10亿次
func TestMulti(t *testing.T){
	//协程的个数
	n:=2
	//每个协程要循环的次数
	lc:=LOOP_COUNT/n
	runtime.GOMAXPROCS(runtime.NumCPU())
	chs:=make([]chan int,n)
	for i:=0;i<n;i++{
		chs[i]=make(chan int)
		c:=chs[i]
		go func(){
			for j:=0;j<lc;j++{
					
			}
			c<-1
		}()
	}
	for _,c:=range(chs){
		<-c
	}
}

//单协程循环10亿次
func TestSingle(t *testing.T){
	for i:=0;i<LOOP_COUNT;i++{
			
	}
}