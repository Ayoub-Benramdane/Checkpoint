package main

import (
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	nb1 := aToi(os.Args[1])
	nb2 := aToi(os.Args[3])
	sign := os.Args[2]
	res := 0
	if (nb1 == 0 && os.Args[1] != "0") || (nb2 == 0 && os.Args[3] != "0") {
		return
	}
	if sign == "+" {
		if hunAdd(nb1, nb2) {
			return
		}
		res = nb1 + nb2
	} else if sign == "-" {
		if hunSub(nb1, nb2) {
			return
		}
		res = nb1 - nb2
	} else if sign == "/" {
		if nb2 == 0 {
			os.Stdout.WriteString("No division by 0")
			return
		}
		res = nb1 / nb2
		if res*nb2 != nb1 {
			return
		}
	} else if sign == "%" {
		if nb2 == 0 {
			os.Stdout.WriteString("No modulo by 0")
			return
		}
		res = nb1 % nb2
	} else if sign == "*" {
		res = nb1 * nb2
		if res/nb1 != nb2 {
			return
		}
	} else {
		return
	}
	os.Stdout.WriteString(iToa(res))
}

func aToi(str string) int {
	res := 0
	t := 1
	for i := 0; i < len(str); i++ {
		if i == 0 && (str[i] == '-' || str[i] == '+') {
			if str[i] == '-' {
				t = -1
			}
			continue
		}
		if str[i] >= '0' && str[i] <= '9' {
			res = res*10 + int(str[i]-'0')
		} else {
			return 0
		}
	}
	return res * t
}

func iToa(nb int) string {
	res := ""
	sign := ""
	if nb < 0 {
		sign = "-"
		nb = -nb
	}
	if nb == 0 {
		return "0"
	}
	for nb != 0 {
		res = string(nb%10+'0') + res
		nb /= 10
	}
	return sign + res
}

func hunAdd(nb1, nb2 int) bool {
	if nb2 > 0 && (nb1 > 9223372036854775807-nb2 || nb2 > 9223372036854775807-nb1) {
		return true
	}
	if nb2 < 0 && (nb1 < -9223372036854775808-nb2 || nb2 < -9223372036854775808-nb1) {
		return true
	}
	return false
}

func hunSub(nb1, nb2 int) bool {
	if nb2 > 0 && (nb1 > 9223372036854775807-nb2 || nb2 > 9223372036854775807-nb1) {
		return true
	}
	if nb2 < 0 && (nb1 < -9223372036854775808-nb2 || nb2 < -9223372036854775808-nb1) {
		return true
	}
	return false
}
