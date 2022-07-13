package main

import (
	"fmt"
	"time"
)

func main() {
	select {
	case <-time.Tick(0 * time.Second):
		fmt.Printf("aa")
	}

	var source1 = make(chan string, 10)
	var source2 = make(chan string, 10)

	test1 := func() {
		for {
			select {
			case <-time.Tick(1 * time.Second):
				select {
				case source := <-source1:
					fmt.Printf("%s\n", source)
				}
			}

		}
	}

	test2 := func() {
		for {
			select {
			case source := <-source2:
				select {
				case <-time.Tick(1 * time.Second):
					fmt.Printf("%s\n", source)
				}
			}

		}
	}

	for i := 0; i < 3; i++ {
		fmt.Printf("%d\n", i)
		go test1()
		go test2()
	}

	source1 <- "hi"
	source1 <- "hi"
	source1 <- "hi"
	source2 <- "hello"
	source2 <- "hello"
	source2 <- "hello"

	time.Sleep(10 * time.Second)

}
