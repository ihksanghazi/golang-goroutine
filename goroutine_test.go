package golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld();
	fmt.Println("OK");

	time.Sleep(1 * time.Second)
}

func displayNumber(number int)  {
	fmt.Println("Display",number)
}

func TestManyGoroutines(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go displayNumber(i)
	}

	time.Sleep(10 * time.Second)
}