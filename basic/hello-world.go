package main

import (
	"fmt"
	_"math"
	_"net/http"	
)

// func main() {
// 	fmt.Println("Hello, Go World!")
// }

// func sayHello(name string) {
// 	fmt.Println("Hello, ", name)
// }

// func main() {
// 	sayHello("jongbeom")
// }

func add(x, y float64) float64 {
	return x+y
}

func subtract(x, y float64) float64 {
	return x-y
}

func main() {
	fmt.Println("5 + 3 = ", add(5, 3))
	fmt.Println("5 - 3 = ", subtract(5, 3))
}