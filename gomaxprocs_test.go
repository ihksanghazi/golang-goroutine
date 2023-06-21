package golang_goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomaxprocs(t *testing.T) {
	group:=sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func ()  {
			time.Sleep(3*time.Second)
			group.Done()
		}()
	}


	totalCPU:= runtime.NumCPU()
	fmt.Println("totalCPU ",totalCPU)

	totalThread:= runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread ",totalThread)

	totalGoroutine:= runtime.NumGoroutine()
	fmt.Println("Total Goroutine ",totalGoroutine)

	group.Done()
}

func TestChangeThreadNumber(t *testing.T) {
	group:=sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func ()  {
			time.Sleep(3*time.Second)
			group.Done()
		}()
	}


	totalCPU:= runtime.NumCPU()
	fmt.Println("totalCPU ",totalCPU)

	runtime.GOMAXPROCS(20)
	totalThread:= runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread ",totalThread)

	totalGoroutine:= runtime.NumGoroutine()
	fmt.Println("Total Goroutine ",totalGoroutine)

	group.Done()
}