package main

import (
	"fmt"
)

func name() string { //we need to mention the return type when we need to return something (int,string,float etc)
	return "Ace of spades"

}

func main() {
	card := name()
	fmt.Printf(card)

}
