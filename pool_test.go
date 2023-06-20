package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)



func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{}{
			return "New"
		},
	}

	pool.Put("Nur")
	pool.Put("Sandy")
	pool.Put("Ihksan")

	for i := 0; i < 10; i++ {
		go func(){
			data:= pool.Get()
			fmt.Println(data)
			time.Sleep(1*time.Second)
			pool.Put(data)
		}()
	}
	time.Sleep(11*time.Second)
	fmt.Println("Selesai")
}