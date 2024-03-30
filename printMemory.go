package main

import (
	"unicode"

	"github.com/01-edu/z01"
)

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

func PrintMemory(arr [10]byte) {
	base := "0123456789abcdef"

	for i := 0; i < len(arr); i++ {
		div := int(arr[i]) / len(base)
		mod := int(arr[i]) % len(base)
		z01.PrintRune(rune(base[div]))
		z01.PrintRune(rune(base[mod]))
		if i != 3 && i != 7 && i != 9 {
			z01.PrintRune(' ')
		}
		if i == 3 || i == 7 || i == 9 {
			z01.PrintRune('\n')
		}
	}
	for i := 0; i < len(arr); i++ {
		if unicode.IsGraphic(rune(arr[i])) {
			z01.PrintRune(rune(arr[i]))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}
