package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	strChan := make(chan string)

	defer close(strChan)

	go func() {
		time.Sleep(2 * time.Second)
		strChan <- "EXAMPLE"
	}()

	data := <-strChan

	fmt.Println("Data from channel:", data)
}

func TestChannelAsParams(t *testing.T) {
	emailChan := make(chan string)

	username := "agusgaming"

	go generateEmail(username, emailChan)

	emailResult := <-emailChan

	fmt.Println("Email:", emailResult)
}

func generateEmail(username string, channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- fmt.Sprintf("%s@example.com", username)

	fmt.Println("Successfully generate the email")
}
