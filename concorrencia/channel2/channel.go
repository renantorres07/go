package main

import (
	"time"
	"fmt"
)

//Channel (canal) - é a forma de comunicação entre goroutines
// channel é um tipo

func doisTresQuartoVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviando dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go doisTresQuartoVezes(2, c)

	a, b := <-c, <-c // recebendo dados do canal
	fmt.Println("B")
	fmt.Println(a, b)

	fmt.Println(<-c)
}