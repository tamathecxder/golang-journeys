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

	close(emailChan)
}

func generateEmail(username string, channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- fmt.Sprintf("%s@example.com", username)

	fmt.Println("Successfully generate the email")
}

func TestInOutChannel(t *testing.T) {
	addrChan := make(chan string)

	defer close(addrChan)

	go setAddress("Jalan cikunir nomor 20, Ciberegbeg Timur, Sukaranda", addrChan)
	go getAddress(addrChan)

	time.Sleep(5 * time.Second)
}

// Channel: Write Only
func setAddress(address string, addrChan chan<- string) {
	time.Sleep(1 * time.Second)

	addrChan <- address

	fmt.Println("Address is now set!")
}

func getAddress(addrChan <-chan string) {
	address := <-addrChan
	fmt.Println("Your address: ", address)
}
