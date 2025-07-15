package goroutines

import (
	"fmt"
	"strconv"
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

// Channel: Write/Set Only
func setAddress(address string, addrChan chan<- string) {
	time.Sleep(1 * time.Second)

	addrChan <- address

	fmt.Println("Address is now set!")
}

// Channel: Read/Receive Only
func getAddress(addrChan <-chan string) {
	address := <-addrChan
	fmt.Println("Your address: ", address)
}

func numToString(num int, resultChan chan<- string) {
	resultChan <- strconv.Itoa(num)
	fmt.Println("Convert to string successfully!")
}

func TestBufferedChannel(t *testing.T) {
	numChan := make(chan int, 2)
	defer close(numChan)

	numChan <- 1
	numChan <- 34

	fmt.Println(cap(numChan))
	fmt.Println(len(numChan))

	fmt.Println(<-numChan)
	fmt.Println(<-numChan)
	fmt.Println(<-numChan) // deadlock, overcapaity

	fmt.Println("Buffered channel is running!")
}

func TestRangeChannel(t *testing.T) {
	numChan := make(chan string)

	go func() {
		for i := 1; i < 10; i++ {
			numChan <- "Data-" + strconv.Itoa(i)
		}

		close(numChan)
	}()

	for data := range numChan {
		fmt.Println("Retrieve:", data)
	}

	fmt.Println("Complete")
}

func TestSelectChannel(t *testing.T) {
	numToStrChan := make(chan string)
	emailChan := make(chan string)

	totalChan := []chan string{numToStrChan, emailChan}

	defer close(numToStrChan)
	defer close(emailChan)

	go numToString(555666, numToStrChan)
	go generateEmail("skuyskuy", emailChan)

	counter := 0
	for {
		select {
		case email := <-emailChan:
			fmt.Println("EMAIL RESULT:", email)
			counter++
		case converted := <-numToStrChan:
			fmt.Println("CONV RESULT:", converted)
			counter++
		default:
			fmt.Println("Loading...")
		}

		if counter == len(totalChan) {
			break
		}
	}
}
