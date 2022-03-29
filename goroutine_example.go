package main

import (
	"fmt"
	"time"
)

func TestGoRoutine() {

	// Calling Goroutine
	go display("Selamat Datang")

	// Calling normal function
	display("Hari yang Cerah")

}
func display(str string) {
	for w := 0; w < 6; w++ {
		time.Sleep(1 * time.Second)
		fmt.Println(w, " ", str)
	}
}
