package main

import "fmt"

func main() {
	for i := 'z'; i >= 'a'; i-- {
		if i%2 == 0 {
			fmt.Printf(string(i))
		} else {
			fmt.Printf(string(i-32))
		}
	}
	fmt.Println()
}
