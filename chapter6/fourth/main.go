package main

import "fmt"

func zero(xPtr *int) {
	*xPtr = 0      //store "0" in the memory location that xPtr refers to
}

func main() {
	x := 5
	zero(&x)       //pass the memory address of x
	fmt.Println(x) //x is now 0
	
}