package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	Frontseat Passenger
}

func main() {
	casey := Passenger{"Casey", 1, false}
	fmt.Println(casey)

	var (
		bill = Passenger{Name: "Bill", TicketNumber: 2}
		ella = Passenger{Name: "Ella", TicketNumber: 3}
	)
	fmt.Println(bill, ella)

	var heidi Passenger
	heidi.Name = "Heidi"
	heidi.TicketNumber = 4
	fmt.Println(heidi)

	casey.Boarded = true
	bill.Boarded = true

	if bill.Boarded {
		fmt.Println("Bill has boarded the buss")
	}
	if casey.Boarded {
		fmt.Println(casey.Name, "has boarded the buss")
	}

	heidi.Boarded = true
	bus := Bus{heidi}
	fmt.Println(bus)
	fmt.Println(bus.Frontseat.Name, "is sitting in the front seat")
}
