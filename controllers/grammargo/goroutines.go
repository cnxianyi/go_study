package grammargo

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

/*
goroutine

 1. 每个并发的执行单元 称作一个 goroutine
 2. go程序启动时,main函数就会在一个单独的goroutine中运行. 即main goroutine
 3. Goroutines是非阻塞的. 函数执行并不会等待Goroutines执行完成
 4. Goroutines会和main goroutine 并发执行
*/
func GoroutinesTest(c *gin.Context) {
	//go spinner(200 * time.Millisecond)
	println(1)
	go say(2)
	go say(3)
	println(4)

	// 创建 channels
	ch1 := make(chan string)
	go sayChannels1(ch1)
	ch1 <- "Ilya"

	go SyncWaitGroup()
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

/*
Channels

1. Channels是 Go中用于在Goroutines之间传递数据的通信机制
2. 允许多个Goroutines共享信息. 且无需锁
*/
func ChannelsTest() {
	println(1)
}

func say(i int) {
	println(i)
}

func sayChannels1(c chan string) {
	s := <-c
	fmt.Printf("hello %s", s)
}

/*
sync.WaitGroup

sync.WaitGroup 等待 所有goroutine完成
*/
func SyncWaitGroup() {
	var wg sync.WaitGroup
	println("-------")
	println(6)
	wg.Add(1) // 增加任务计数.表示有一个任务需要等待
	go func() {
		defer wg.Done() // 任务执行,-1任务
		time.Sleep(2 * time.Second)
		say(7)
	}()

	wg.Wait() // 等待wg任务计数归零. 即所有任务完成

	wg.Add(1)
	go func() {
		defer wg.Done()
		say(8)
	}()

	wg.Wait()

	println(9)
}

// func spinner(delay time.Duration) {
// 	for {
// 		for _, r := range `-\|/` {
// 			fmt.Printf("\r%c", r)
// 			time.Sleep(delay)
// 		}
// 	}
// }
