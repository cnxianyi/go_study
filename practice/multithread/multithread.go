package multithread

import (
	"fmt"
	"sync"
)

// 目标: 使用两个线程顺序打印0~20

func PrintWithTwoThread() {
	var wg sync.WaitGroup

	wg.Add(2)

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer wg.Done()
		for num := range ch1 {
			if num > 20 {
				close(ch2)
				return
			}

			fmt.Printf("进程1: %d\n", num)
			ch2 <- num + 1
		}
	}()

	go func() {
		defer wg.Done()
		for num := range ch2 {
			if num > 20 {
				close(ch1)
				return
			}

			fmt.Printf("进程2: %d\n", num)
			ch1 <- num + 1
		}
	}()

	ch1 <- 0

	wg.Wait()
}
