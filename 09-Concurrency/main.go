package main

import (
	"fmt"
	"time"
)

/*
 A goroutine is a lightweight thread managed by the Go runtime.
go f(x, y, z)
*/

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

/*
Channels are a typed conduit through which you can send and receive values with
the channel operator,  <-.

E.g
ch <- v    [Send v to channel ch.]
v := <-ch  [Receive from ch, and]
           [assign value to v.]

the data flows in the direction of the arrow
*/

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum = sum + v
	}
	c <- sum // send sum to c
}

/*
Buffered Channels
Channels can be buffered.
Provide the buffer length as the second argument to make to initialize a buffered channel:
{ch := make(chan int, 100)}
*/
//------------------------------
/*
Range and Close:
A sender can close a channel to indicate that no more values will be sent.
Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after
{v, ok := <-ch}
ok is false if there are no more values to receive and the channel is closed.

Note:
Only the sender should close a channel, never the receiver.
Sending on a closed channel will cause a panic.

Another note:
Channels aren't like files; you don't usually need to close them.
Closing is only necessary when the receiver must be told there are no more values coming,
such as to terminate a range loop.
*/
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

//----------------------
/*
	Select
The select statement lets a goroutine wait on multiple communication operations.
A select blocks until one of its cases can run, then it executes that case.
It chooses one at random if multiple are ready.
*/

func main() {
	//  Goroutines
	//fmt.Println("Goroutines")
	go say("World")
	say("hello")

	// Channels
	s := []int{7, 2, 8, -9, 4, 0}
	// Like maps and slices, channels must be created before use: {ch := make(chan int)}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c
	//fmt.Println("Channels")
	fmt.Println(x, y, x+y)

	//Buffered Channels
	Bch := make(chan int, 3)
	Bch <- 1
	Bch <- 2
	Bch <- 3
	//fmt.Println("Buffered Channels")
	fmt.Println(<-Bch)
	fmt.Println(<-Bch)
	fmt.Println(<-Bch)

	// Range and Close
	RC := make(chan int, 5)
	// Cap of channel is the buffer size
	go fibonacci(cap(RC), RC)
	//fmt.Println("Range and Close")
	for i := range RC {
		fmt.Println(i)
	}

	/*
		channels are great for communication among goroutines.
		This concept is called mutual exclusion,
		and the conventional name for the data structure that provides it is MUTEX.

		Go's standard library provides mutual exclusion with sync.Mutex and its two methods:
		-Lock
		-Unlock
	*/
}
