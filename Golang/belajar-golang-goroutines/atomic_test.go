package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		go func() {
			group.Add(1)
			for j := 1; j <= 100; j++ {
				// x = x + 1
				atomic.AddInt64(&x, 1)
			}
			group.Done()
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x)
}