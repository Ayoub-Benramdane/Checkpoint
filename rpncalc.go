package main

import (
	"fmt"
	"os"
)

func aTOi(str string) int {
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

func main() {
	if len(os.Args[1:]) != 1 {
		fmt.Println("Error")
		return
	}
	arg := os.Args[1]
	nbr := []int{}
	res := 0
	val := 0
	for i := 0; i < len(arg); i++ {
		if !(arg[i] >= '0' && arg[i] <= '9') && arg[i] != ' ' {
			if len(nbr) == 2 {
				if arg[i] == '+' || arg[i] == '-' || arg[i] == '*' || arg[i] == '/' || arg[i] == '%' {
					if arg[i] == '+' {
						res = nbr[0] + nbr[1]
						nbr = nil
						nbr = append(nbr, res)
					} else if arg[i] == '-' {
						res = nbr[0] - nbr[1]
						nbr = nil
						nbr = append(nbr, res)
					} else if arg[i] == '/' {
						if arg[i-1] == 0 {
							fmt.Println("No division by 0")
						}
						res = nbr[0] / nbr[1]
						nbr = nil
						nbr = append(nbr, res)
					} else if arg[i] == '%' {
						if arg[i-1] == 0 {
							fmt.Println("No modulo by 0")
						}
						res = nbr[0] % nbr[1]
						nbr = nil
						nbr = append(nbr, res)
					} else if arg[i] == '*' {
						res = nbr[0] * nbr[1]
						nbr = nil
						nbr = append(nbr, res)
					}
				}
			} else if len(nbr) == 3 {
				if arg[i] == '+' || arg[i] == '-' || arg[i] == '*' || arg[i] == '/' || arg[i] == '%' {
					if arg[i] == '+' {
						res = nbr[0] + nbr[1]
						val = nbr[2]
						nbr = nil
						nbr = append(nbr, res)
						nbr = append(nbr, val)
					} else if arg[i] == '-' {
						res = nbr[0] - nbr[1]
						val = nbr[2]
						nbr = nil
						nbr = append(nbr, res)
						nbr = append(nbr, val)
					} else if arg[i] == '/' {
						if arg[i-1] == 0 {
							fmt.Println("No division by 0")
						}
						res = nbr[0] / nbr[1]
						val = nbr[2]
						nbr = nil
						nbr = append(nbr, res)
						nbr = append(nbr, val)
					} else if arg[i] == '%' {
						if arg[i-1] == 0 {
							fmt.Println("No modulo by 0")
						}
						res = nbr[0] % nbr[1]
						val = nbr[2]
						nbr = nil
						nbr = append(nbr, res)
						nbr = append(nbr, val)
					} else if arg[i] == '*' {
						res = nbr[0] * nbr[1]
						val = nbr[2]
						nbr = nil
						nbr = append(nbr, res)
						nbr = append(nbr, val)
					}
				}
			} else {
				fmt.Println("Error")
				return
			}
		} else if arg[i] != ' ' {
			nbr = append(nbr, aTOi(string(arg[i])))
		}
	}
	fmt.Println(res)
}
