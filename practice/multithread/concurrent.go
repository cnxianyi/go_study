package multithread

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"sync"
)

// 并发测试: 并发五个任务
// 并捕获第一个错误
func ConcurrentTestA() {

	var wg sync.WaitGroup
	var firstError error
	var mu sync.Mutex

	runTask := func(task func() error) {
		defer wg.Done()

		if err := task(); err != nil {
			mu.Lock()
			if firstError == nil {
				firstError = err
			}
			mu.Unlock()
		}
	}

	wg.Add(5)

	go runTask(task)
	go runTask(task)
	go runTask(task)
	go runTask(task)
	go runTask(task)

	wg.Wait()

	if firstError != nil {
		fmt.Println(firstError)
	}
}

func task() error {
	r, _ := rand.Int(rand.Reader, big.NewInt(100))
	fmt.Println(r)
	r1, _ := strconv.Atoi(r.String())

	if r1 < 20 {
		return fmt.Errorf("small")
	}

	return nil
}
