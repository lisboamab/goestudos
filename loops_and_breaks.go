package main

import "fmt"

func main() {
	for i := 0; i <= 20.0; i++ {
		if i%2 != 0 {
			fmt.Printf("%v\t%T\timpar\n", i, i)
		} else {
			fmt.Printf("%v\t%T\tpar\n", i, i)
		}
	}
}
