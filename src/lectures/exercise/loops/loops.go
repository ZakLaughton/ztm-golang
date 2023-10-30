//--Summary:
//  Implement the classic "FizzBuzz" problem using a `for` loop.
//
//--Requirements:
//* Print integers 1 to 50, except:
//  - Print "Fizz" if the integer is divisible by 3
//  - Print "Buzz" if the integer is divisible by 5
//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
//
//--Notes:
//* The remainder operator (%) can be used to determine divisibility

package main

func main() {
	myNum := 0

	for i := 1; i <= 50; i++ {
		myNum = i
		if myNum%3 == 0 && myNum%5 == 0 {
			println("FizzBuzz")
		} else if myNum%3 == 0 {
			println("Fizz")
		} else if myNum%5 == 0 {
			println("Buzz")
		} else {
			println(myNum)
		}
	}
}
