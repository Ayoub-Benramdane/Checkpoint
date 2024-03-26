package main

func FifthAndSkip(str string) string {
	var s string
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			s += string(str[i])
		}
	}
	r := []byte(s)
	for i := 0; i < len(r); i++ {
		if r[i] <= '9' && r[i] >= '0' {
			return "Invalid Input" + "\n"
		}
		if i%6 == 5 {
			r[i] = ' '

		}
	}
	return string(r) + "\n"
}
