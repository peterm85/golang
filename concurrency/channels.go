package main

import (
	"fmt"
	"time"
)

/*
By default, there’s no storage capacity in channels,
which means the message must flow from the sender to the receiver immediately.
If there’s no receiver, a message is stuck with the sender.

*/
func deadLockDataChannel() {
	dataChannel := make(chan string)
	dataChannel <- "Some Sample Data"
	fmt.Println(<-dataChannel)
}

/*
By adding a buffer to the channel, we give it a capacity to save some messages internally,
allowing the sender to proceed with its work even if the data is not extracted on the other end
and removing the deadlock.
*/
func bufferedDataChannel() {
	dataChannel := make(chan string, 3)
	dataChannel <- "Some Sample Data"
	dataChannel <- "Some Other Sample Data"
	dataChannel <- "Buffered Channel"
	fmt.Println(<-dataChannel)
	fmt.Println(<-dataChannel)
	fmt.Println(<-dataChannel)
}

/*
Listen messages from a subscriber chan
*/
func subscribe(name string, subscriber <-chan string) {
	for {
		select {
		case message := <-subscriber:
			fmt.Printf("%q: %q\n", name, message)
		case <-time.After(time.Millisecond):
		}
	}
}

/*
Send a message from publisher channel to all his associated subscribers channels
*/
func publish(subscriptions map[chan string][]chan string) {
	for {
		for publisher, subscribers := range subscriptions {
			select {
			case message := <-publisher:
				for _, subscriber := range subscribers {
					subscriber <- message
				}
			case <-time.After(time.Millisecond):
			}
		}
	}
}

func pubSubChannel() {

	sub1 := make(chan string)
	sub2 := make(chan string)
	sub3 := make(chan string)

	go subscribe("Subscriber 1", sub1)
	go subscribe("Subscriber 2", sub2)
	go subscribe("Subscriber 3", sub3)

	pub1 := make(chan string)
	pub2 := make(chan string)

	subscriptions := map[chan string][]chan string{
		pub1: {sub1, sub2},
		pub2: {sub2, sub3},
	}

	go publish(subscriptions)

	//Let's send messages!
	pub1 <- "Hello, World!"
	pub2 <- "Hi, Universe!"
	pub1 <- "Goodbye, Cruel World!"

	fmt.Printf("\nAppending Subscriber 3 to pub1:\n\n")
	subscriptions[pub1] = append(subscriptions[pub1], sub3)
	pub1 <- "Just kidding!"

	time.Sleep(2 * time.Second) //Some race conditions was found
}

/*func main() {
	arg := os.Args[1]

	switch arg {
	case "1":
		deadLockDataChannel()
	case "2":
		bufferedDataChannel()
	case "3":
		pubSubChannel()
	}
}*/
