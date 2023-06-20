package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func onlyOne(){
	counter++
}

func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func(){
			once.Do(onlyOne)
			group.Done()
		}()	
	}
	group.Wait()
	fmt.Println("counter ",counter)
}