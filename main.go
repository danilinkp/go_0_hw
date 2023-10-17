package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Printf("My favourite number is %d", rand.Int())

}

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}
