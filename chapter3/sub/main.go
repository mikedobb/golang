package main

import "fmt"

func main(){
    fmt.Println("Enter a Fahrenheit temp: ")
    var input float64
    fmt.Scanf("%f", &input)
    
    input -= 32

    temp := input * 5

    output := temp / 9

    fmt.Println(output)
}
