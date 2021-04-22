package errhandling

import (
	"fmt"
	"time"
)

func validation(input string, dlq chan string) {
	fmt.Println("Validating max length...")
	if len(input) > 10 {
		dlq <- "param [input] length must be lower than 10"
	}

	fmt.Println("Validating min length...")
	if len(input) < 2 {
		dlq <- "param [input] length must be higher than 2"
	}

	fmt.Println("Validating customized value...")
	if input == "Hello!" {
		dlq <- "param [input] length must not be " + input
	}
}

func process(input string, dlq chan string) {
	validation(input, dlq)
	//Doing stuffs
	time.Sleep(2 * time.Second)
}

func monitor(dlq <-chan string) {
	for {
		select {
		case message := <-dlq:
			fmt.Println(message)
		case <-time.After(time.Millisecond):
		}
	}
}

func RunWorkspace(input string) {
	dlq := make(chan string)
	go monitor(dlq)
	go process(input, dlq)
	time.Sleep(5 * time.Second)
}
