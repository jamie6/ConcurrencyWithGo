package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	randBoring("hey, you up?")
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