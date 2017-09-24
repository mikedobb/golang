package main

import "fmt"

func main(){
    var x string = ("Goodbye, Cruel World")
    fmt.Println(x)
    const y string = "that was easy"
    fmt.Println(y)
    x += y
    fmt.Println(x)

}
