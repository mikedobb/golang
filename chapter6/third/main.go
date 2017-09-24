package main

import "fmt"

func main() {
    defer func() {
    	str := recover()   //recover is deferred to ensure that is runs
    	fmt.Println(str)
    }()
    panic("PANIC")
}