/**
* 1. go的协程goroutine是并发的能力，类似于js的Promise, Async/await，并不是并行；
* 2. go的协程允许使用同步的方式写异步代码；
* 3. 协程一般与通道channel配合传输数据，使用 <- 符号传输；
* 4. 执行代码时，会先同步执行goroutine，直到遇到包含 <- 读取或写入通道的语句，就会等待异步代码，
* 也就是协程执行完毕拿到通道数据才会继续向下运行；
* 5. 只有给通道发送数据的协程需要关闭通道，接收者永远不会需要，关闭使用defer close(chan)这样的方式；
*/

package torrent

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Decode() {
	ch := make(chan string)
	fmt.Println("hello 1")
	go request(ch)
	fmt.Println("hello 2")
	fmt.Println(<-ch)
	fmt.Println("hello 3")
}

func request(ch chan string) {
	resp, _ := http.Get("https://www.baidu.com")
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("end")
	ch <- string(data)
	defer close(ch)
}

/*
  输出结果（说明了协程代码执行的顺序）
  hello 1
  hello 2
  end
  <html>。。。</html>
  hello 3
*/
