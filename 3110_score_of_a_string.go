package main

// fmt package provides the function to print anything
import (
	"fmt"
)

func main() {
	s := "hello"
	lens := len(s)
	summs := 0
	prom := 0
	for i := 0; i < lens-1; i++ {
		prom = int(s[i] - s[i+1])
		if prom < 0 {
			prom = prom * (-1)
		}
		summs = summs + prom
	}
	fmt.Println("(Printing the ASCII value using specifier)")
}
