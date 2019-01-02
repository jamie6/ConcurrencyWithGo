// select statement to control multiple channels
// the time.After function returns a channel that blocks for the specified duration.
// After the interval, the channel delivers the current time, once

// when using timeout,
// create the timer once, outside the loop, to time out the entire conversation.
// (in previous, we had a timeout for each message.)
package main
import "fmt"
import "time"
import "math/rand"

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		select {
		case s := <-input1: c <- s
		case s := <-input2: c <- s
		}
	}()
	return c
}

func boring(msg string) <-chan string{ // Returns receive-only channel of strings.
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function.
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // Return the channel to the caller.
}

func main() {
	c := boring("Joe")
timeout := time.After(5*time.Second)
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
		//case <-time.After(1*time.Second):
			fmt.Println("You're too slow.")
			return
		}
	}
}