//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Character struct {
	Health    int
	MaxHealth int
	Energy    int
	MaxEnergy int
	Name      string
}

func (c *Character) heal() {
	fmt.Println("Health:", c.Health, "Healing...")
	c.Health += 1
	fmt.Println("New health:", c.Health)
}

func (c *Character) energize() {
	fmt.Println("Energy:", c.Energy, "Energizing...")
	c.Energy += 1
	fmt.Println("New energy:", c.Energy)
}

func main() {
	timmy := Character{
		Health:    10,
		MaxHealth: 20,
		Energy:    10,
		MaxEnergy: 20,
	}
	fmt.Println("Initial:", timmy)
	timmy.heal()
	timmy.energize()
}
