package main

import "fmt"
import "time"

func countdown() {

	//Simple loop
	fmt.Println("When is it?")
	for i := 10; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
	fmt.Println("Happy New Year")
}
