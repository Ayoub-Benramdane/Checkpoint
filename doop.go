package main

import (
	"fmt"
	"os"
)

func aToi(str string) int {
	res := 0
	t := 1
	for i := 0; i < len(str); i++ {
		if str[i] == '-' && i == 0 {
			t = -1
		}
		if str[i] >= '0' && str[i] <= '9' {
			res = res*10 + int(str[i]-'0')
		}
	}
	return res * t
}

func iToa(nb int) string {
	res := ""
	for nb != 0 {
		res = string(nb%10+'0') + res
		nb /= 10
	}
	return res
}

func main() {
	nb1 := aToi(os.Args[1])
	nb2 := aToi(os.Args[3])
	signe := os.Args[2]
	if !((nb1 >= 0 && nb1 <= 9) || (nb2 >= 0 && nb2 <= 9)) {
		fmt.Println("")
	}
	if signe == "+" {
		fmt.Println(iToa(nb1 + nb2))
	} else if signe == "-" {
		fmt.Println(iToa(nb1 - nb2))
	} else if signe == "/" {
		if nb2 == 0 {
			fmt.Println("No division by 0")
		}
		fmt.Println(iToa(nb1 / nb2))
	} else if signe == "%" {
		if nb2 == 0 {
			fmt.Println("No modulo by 0")
		}
		fmt.Println(iToa(nb1 % nb2))
	} else if signe == "*" {
		fmt.Println(iToa(nb1 * nb2))
	} else {
		fmt.Println("")
	}
}
