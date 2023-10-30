package main

import "fmt"

type Space struct {
	occupied bool
}

type ParkingLot struct {
	spaces []Space
}

func occupySpace(lot *ParkingLot, spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLot) occupySpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLot) vacateSpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = false
}

func main() {
	lot := ParkingLot{spaces: make([]Space, 4)}
	fmt.Println("Initial:", lot)
	lot.occupySpace(1)
	fmt.Println(lot)
	// Same as the receiver above
	occupySpace(&lot, 2)
	fmt.Println(lot)

	lot.vacateSpace(2)
	fmt.Println(lot)
}
