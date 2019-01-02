//concurrent 
package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	go randBoring("hey, you up?")
	fmt.Println("*farts")
	time.Sleep(2*time.Second)
	fmt.Println("Nope.")
}

func boring(msg string){
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Second)
	}
}

func randBoring(msg string){
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}