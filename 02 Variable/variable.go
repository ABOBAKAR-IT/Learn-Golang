package main

import "fmt"

func main() {
	fmt.Println("Variable in Go")

	var a string = "abobakar"                            // String Variable Declaration
	fmt.Printf("The value of variable 'a' is %v \n ", a) // %v is used for printing the value of variable , %v working with Printf not with Println
	fmt.Printf("The type of variable 'a' is %T \n", a)
	fmt.Println("The value of constant 'c1' is %v \n ", a)
	// Implicitly Type Inference
	var b = "abobakar"
	fmt.Printf("The value of variable 'b' is %v \n ", b)
	fmt.Printf("The type of variable 'b' is %T \n", b)
	//b = 123 // error: cannot use 123 (type int) as type string in assignment

	const c string = "Abdulbasit" // Constant Variable Declaration
	fmt.Printf("The value of constant 'c' is %v \n ", c)

	var d, e int = 123, 132 // Multiple Variable Declaration
	fmt.Printf("The value of variable 'd' is %v \n ", d)
	fmt.Printf("The value of variable 'e' is %v \n ", e)

	var f, g int
	fmt.Printf("Value of f is %v  and Value of g is %v \n", f, g) // Both are zero values

	var h, i = 123, "Abdulbasit" // Shorthand Assignment
	fmt.Printf("Value of h is %v  and Value of i is %v \n", h, i)

	sum := 125 + 125 // Shorthand Assignment with implicit declaration
	fmt.Printf("Value of sum is %v \n", sum)

	var pi float64 = 14.123412
	fmt.Printf("Value of pi is %v \n", pi)

	var check bool = true
	fmt.Printf("Value of check is %v \n", check)

}
