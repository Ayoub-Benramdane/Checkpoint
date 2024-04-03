package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args[1:]) != 1 {
		return
	}
	arg, err := strconv.Atoi(os.Args[1])
	res := ""
	res1 := ""
	resFinal := ""
	if err != nil || arg >= 4000 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}
	for arg > 0 {
		if arg >= 1000 {
			res += "M"
			res1 += "M+"
			arg -= 1000
		} else if arg >= 900 {
			res += "CM"
			res1 += "(M-C)+"
			arg -= 900
		} else if arg >= 500 {
			res += "D"
			res1 += "D+"
			arg -= 500
		} else if arg >= 400 {
			res += "CD"
			res1 += "(D-C)+"
			arg -= 400
		} else if arg >= 100 {
			res += "C"
			res1 += "C+"
			arg -= 100
		} else if arg >= 90 {
			res += "XC"
			res1 += "(C-X)+"
			arg -= 90
		} else if arg >= 50 {
			res += "L"
			res1 += "L+"
			arg -= 50
		} else if arg >= 40 {
			res += "XL"
			res1 += "(L-X)+"
			arg -= 40
		} else if arg >= 10 {
			res += "X"
			res1 += "X+"
			arg -= 10
		} else if arg >= 9 {
			res += "IX"
			res1 += "(X-I)+"
			arg -= 9
		} else if arg >= 5 {
			res += "V"
			res1 += "V+"
			arg -= 5
		} else if arg >= 4 {
			res += "IV"
			res1 += "(V-I)+"
			arg -= 4
		} else if arg >= 1 {
			res += "I"
			res1 += "I+"
			arg -= 1
		}
	}
	for i := 0; i < len(res1)-1; i++ {
		resFinal += string(res1[i])
	}
	fmt.Println(resFinal)
	fmt.Println(res)
}
