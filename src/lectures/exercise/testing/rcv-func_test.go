//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests

package main

import "testing"

// * Write unit tests that ensure:
//   - Health & energy can not go above their maximums
func TestHeal(t *testing.T) {
	timmy := Character{
		Health:    20,
		MaxHealth: 20,
	}
	timmy.heal()
	if timmy.Health > timmy.MaxHealth {
		t.Errorf("Health should not be greater than MaxHealth")
	}
}

func TestEnergize(t *testing.T) {
	timmy := Character{
		Energy:    20,
		MaxEnergy: 20,
	}
	timmy.energize()
	if timmy.Energy > timmy.MaxEnergy {
		t.Errorf("Energy should not be greater than MaxEnergy")
	}
}

//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
