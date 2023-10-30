//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

// * Create a structure to store items and their security tag state
//   - Security tags have two states: active (true) and inactive (false)
type Item struct {
	name string
	tag  SecurityTag
}

// * Create functions to activate and deactivate security tags using pointers
func activateTag(tag *SecurityTag) {
	*tag = Active
}

func deactivateTag(tag *SecurityTag) {
	*tag = Inactive
}

// * Create a checkout() function which can deactivate all tags in a slice
func checkout(items []Item) {
	fmt.Println("checking out...")
	for i := range items {
		deactivateTag(&items[i].tag)
	}
}

const (
	Active   = true
	Inactive = false
)

type SecurityTag bool

func main() {
	// * Perform the following:
	//   - Create at least 4 items, all with active security tags
	shirt := Item{name: "shirt", tag: Active}
	pants := Item{name: "pants", tag: Active}
	shoes := Item{name: "shoes", tag: Active}
	hat := Item{name: "hat", tag: Active}
	//   - Store them in a slice or array
	items := []Item{shirt, pants, shoes, hat}
	fmt.Println(items)

	//   - Deactivate any one security tag in the array/slice
	deactivateTag(&items[0].tag)
	fmt.Println(items)

	//   - Call the checkout() function to deactivate all tags
	checkout(items)
	//   - Print out the array/slice after each change
	fmt.Println(items)

}
