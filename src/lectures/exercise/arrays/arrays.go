//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name  string
	price float32
}

func main() {
	shoppingList := [4]Product{
		{name: `milk`, price: 1.99},
		{name: `eggs`, price: 2.99},
		{name: `bread`, price: 3.99},
	}

	fmt.Println("Last item:", shoppingList[len(shoppingList)-1])
	fmt.Println("Total number of items:", len(shoppingList))
	fmt.Println("Total cost:", shoppingList[0].price+shoppingList[1].price+shoppingList[2].price)

	shoppingList[3] = Product{name: `cheese`, price: 4.99}

	fmt.Println("Last item:", shoppingList[len(shoppingList)-1])
	fmt.Println("Total number of items:", len(shoppingList))
	fmt.Println("Total cost:", shoppingList[0].price+shoppingList[1].price+shoppingList[2].price)

	//   - Print to the terminal:
	//   - The total number of items
	//   - The total cost of the items
	//   - Add a fourth product to the list and print out the
	//     information again
}
