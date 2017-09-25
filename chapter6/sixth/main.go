package main

import "fmt"

func halved(x uint) (ret uint, eveness bool){
	ret = x / 2
	eveness = false
	if x % 2 == 0 {
		eveness = true

	}
	return ret, eveness

}

func muhammedAli(args ...int) int {
	greatest := 0
	for _, v := range args {
		if v >= greatest {
			greatest = v
		}
	}
	return greatest
}

func makeOddGenerator() func() uint {   
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i +=2
		return
	}
}

func fibonacci(n int) int {
	switch n{
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fibonacci(n-1) + fibonacci(n-2)
	}
}




func main() {
        genfunc1 := makeOddGenerator()
        genfunc2 := makeOddGenerator()

        fmt.Println(genfunc1())
        fmt.Println(genfunc2())
        fmt.Println(genfunc1())
        fmt.Println(genfunc2())
        fmt.Println(fibonacci(8))

	fmt.Println(halved(10))
	fmt.Println(halved(21))
	fmt.Println(muhammedAli(23, 51, 5115, 50389547, 403928,583748))
    
}
