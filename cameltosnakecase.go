package main

import (
	"fmt"
)

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}

func CamelToSnakeCase(s string) string {
	str := ""
	isValid := false
	if !(s[len(s)-1] >= 'A' && s[len(s)-1] <= 'Z') {
		isValid = true
	}
	for i, c := range s {
		if isValid && i != 0 && c >= 'A' && c <= 'Z' {
			str += "_"
		}
		str += string(c)
	}
	return str
}
