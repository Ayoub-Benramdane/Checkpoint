package main

import "fmt"

func main() {
	table := []int{1, 2, 3}
	ac := 93
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	fmt.Println()

	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}

func FoldInt(f func(int, int) int, a []int, n int) {
	for i := 0; i < len(a); i++ {
		n = f(a[i], n)
	}
	fmt.Println(n)
}

func Add(nb, nb2 int) int {
	res := 0
	res = nb + nb2
	return res
}

func Mul(nb, nb2 int) int {
	res := 0
	res = nb * nb2
	return res
}

func Sub(nb, nb2 int) int {
	res := 0
	res = nb2 - nb
	return res
}
