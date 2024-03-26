package main

func aToi(str string) int { //125
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

func iToa(nb int) string { //125
	res := ""
	for nb != 0 {
		res = string(nb%10+'0') + res
		nb /= 10
	}
	return res
}

func doop(args []string) string {
	nb1 := aToi(args[1])
	nb2 := aToi(args[3])
	signe := args[2]
	if !((nb1 >= 0 && nb1 <= 9) || (nb2 >= 0 && nb2 <= 9)) {
		return ""
	}
	if signe == "+" {
		return iToa(nb1 + nb2)
	} else if signe == "-" {
		return iToa(nb1 - nb2)
	} else if signe == "/" {
		if nb2 == 0 {
			return "No division by 0"
		}
		return iToa(nb1 / nb2)
	} else if signe == "%" {
		if nb2 == 0 {
			return "No modulo by 0"
		}
		return iToa(nb1 % nb2)
	} else if signe == "*" {
		return iToa(nb1 * nb2)
	} else {
		return ""
	}
}
