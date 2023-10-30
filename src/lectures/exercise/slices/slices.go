//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

func main() {
	var assemblyLine [3]Part
	fmt.Println("assemblyLine:", assemblyLine)
	assemblyLine[0] = "Part 1"
	fmt.Println("assemblyLine:", assemblyLine)
	assemblyLine[1] = "Part 2"
	fmt.Println("assemblyLine:", assemblyLine)

	firstTwoParts := assemblyLine[:2]
	fmt.Println("assemblyLine:", assemblyLine)
	fmt.Println("slice:", firstTwoParts)
}
