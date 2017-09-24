package main

import "fmt"

func f() (int, int) {
	return 5, 6  // Reminder that Go funcitons can return multiple values
}

func add(args ...int) int {
	total := 0
    for _, v := range args {
    	total += v
    }
    return total

}

func main() {
	x, y := f()

	fmt.Println(x, y)
	fmt.Println(add(1,2,3,4,5,6,7,8,9,10))
}