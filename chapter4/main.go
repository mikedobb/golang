package main

import "fmt"

func main(){
    oddity := "odd"
    name := "zero"
    for i:= 1; i <= 13; i ++ {
        if i % 2 == 0{
            oddity = "even"
        } else {
            oddity = "odd"
        }

        switch i {
            case 0: name = "zero"
            case 1: name = "one"
            case 2: name = "two"
            case 3: name = "three"
            case 4: name = "four"
            case 5: name = "five"
            case 6: name = "six"
            case 7: name = "seven"
            case 8: name = "eight"
            case 9: name = "nine"
            case 10: name = "ten"
            default: name = "Unknown Number"   
        }
        fmt.Println(name, " is ", oddity)
    }
}
