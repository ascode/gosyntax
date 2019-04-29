// _Channels_ are the pipes that connect concurrent
// goroutines. You can send values into channels from one
// goroutine and receive those values into another
// goroutine.

package main

import "fmt"

var messages = make(chan string)
var counter = make(chan int)
var countnum = 0
var i = 0

func main() {

	// Create a new channel with `make(chan val-type)`.
	// Channels are typed by the values they convey.


	// _Send_ a value into a channel using the `channel <-`
	// syntax. Here we send `"ping"`  to the `messages`
	// channel we made above, from a new goroutine.
	for ; i < 15; i++ {
		go stackMessage(i)
	}

	// The `<-channel` syntax _receives_ a value from the
	// channel. Here we'll receive the `"ping"` message
	// we sent above and print it out.
	///for len(count) < 15 {
	//	time.Sleep(1000)
	//	println(len(messages))
	for countnum < 15 {
		j := <- counter
		countnum += j
		fmt.Println(countnum)
		for _ = range messages {
			msg := <- messages
			fmt.Println(msg)
		}

	}


	//	countfunc()
	//}

}

func stackMessage(is int) {
	messages <- "ping" + fmt.Sprint(is)
	counter <- 1
}
