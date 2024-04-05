package main

import (
	"fmt"
)

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, 1))
	fmt.Printf("%#v\n", Slice(a, 2, 4))
	fmt.Printf("%#v\n", Slice(a, -3))
	fmt.Printf("%#v\n", Slice(a, -2, -1))
	fmt.Printf("%#v\n", Slice(a, 2, 0))
}

func Slice(a []string, nbrs ...int) []string {
	res := []string{}
	for i := 0; i < len(a); i++ {
		if len(nbrs) == 1 {
			if nbrs[0] > 0 {
				if i >= nbrs[0] && i < len(a) {
					res = append(res, a[i])
				}
			} else if nbrs[0] < 0 {
				if i >= len(a)+nbrs[0] && i < len(a) {
					res = append(res, a[i])
				}
			}
		} else {
			if nbrs[1] > nbrs[0] {
				if nbrs[0] > 0 {
					if i >= nbrs[0] && i < nbrs[1] {
						res = append(res, a[i])
					}
				} else {
					if i >= len(a)+nbrs[0] && i < len(a)+nbrs[1] {
						res = append(res, a[i])
					}
				}
			} else {
				return nil
			}
		}
	}
	return res
}
