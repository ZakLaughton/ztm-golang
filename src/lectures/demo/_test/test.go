package main

import "fmt"

type Player struct {
	health   int
	position int
}

func main() {
	myVar := Player{10, 20}
	vari, ok := myVar.(Player)
	fmt.Println(
		"vari:", vari,
		"ok:", ok,
	)
}
