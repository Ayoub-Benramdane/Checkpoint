package main

import "fmt"

func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}

func IsCapitalized(s string) bool {
	g := []rune(s)
	if len(s) == 0 || g[0] >= 'a' && g[0] <= 'z' {
		return false
	}
	for i := 0; i < len(s); i++ {
		if g[i] == ' ' {
			if g[i+1] >= 'a' && g[i+1] <= 'z' {
				return false
			}
		}
	}
	return true
}
