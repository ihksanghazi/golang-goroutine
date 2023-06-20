package golang_goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func ()  {
		time.Sleep(2*time.Second)
		channel <- "Nursandy Ihksan"
		fmt.Println("Selesai Mengirim Data Ke Channel")
	}()

	data := <- channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func giveMeResponse(channel chan string){
	time.Sleep(2 * time.Second)
	channel <- "Nursandy Ihksan"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go giveMeResponse(channel)

	data := <- channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)	
}

func onlyIn(channel chan<- string){
	time.Sleep(2*time.Second)
	channel<- "Nursandy Ihksan"
}

func onlyOut(channel <-chan string){
	data:= <-channel;
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string);

	go onlyIn(channel)
	go onlyOut(channel)

	time.Sleep(3*time.Second)
	close(channel)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string,3)
	defer close(channel)

	go func(){
		channel<-"Nur"
		channel<-"Sandy"
		channel<-"Ihksan"
	}()

	go func(){
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func(){
		for i := 0; i < 10; i++ {
			channel<- "Perulangan ke - " + strconv.Itoa(i)
		}
		close(channel)
		}()

		for data := range channel {
			fmt.Println("Mengirim Data ", data)	
		}

		fmt.Println("Selesai")
	}