package main

func BeZero(slice []int) []int {
	if len(slice) == 0 {
		return []int{}
	}
	for i := 0; i < len(slice); i++ {
		slice[i] = 0
	}
	return slice
}

/*
func BinaryCheck(nbr int) int {
	if nbr%2 == 0 {
		return 1
	}
	return 0
}
*/

/*
const memorySize = 2048

func brainFuck(source string) {
	memory := make([]byte, memorySize)
	ptr := 0
	pc := 0
	loopStack := make([]int, 0)
	output := ""

	for pc < len(source) {
		switch source[pc] {
		case '>':
			ptr = (ptr + 1) % memorySize
		case '<':
			ptr = (ptr - 1 + memorySize) % memorySize
		case '+':
			memory[ptr]++
		case '-':
			memory[ptr]--
		case '.':
			output += string(memory[ptr])
		case '[':
			if memory[ptr] == 0 {
				nested := 1
				for nested > 0 {
					pc++
					if source[pc] == '[' {
						nested++
					} else if source[pc] == ']' {
						nested--
					}
				}
			} else {
				loopStack = append(loopStack, pc)
			}
		case ']':
			if memory[ptr] != 0 {
				if len(loopStack) > 0 {
					pc = loopStack[len(loopStack)-1] - 1
				}
			} else {
				loopStack = loopStack[:len(loopStack)-1]
			}
		}

		pc++
	}

	fmt.Println(output)
}
*/

/*
func ByeByeFirst(strings []string) []string {
	if len(strings) == 0 {
		return []string{}
	}
	return strings[1:]

}
*/

/*
func capitalizeFirst(input string) string {
	if len(input) == 0 {
		return ""
	}
	firstChar := input[0]
	if firstChar >= 'a' && firstChar <= 'z' {
		capitalized := firstChar - 'a' + 'A'
		return string(capitalized) + input[1:]
	}
	return input
}
func titleCase(input string) string {
	words := make([]string, 0)
	wordStart := 0
	for i := 0; i < len(input); i++ {
		if input[i] == ' ' {
			if i > wordStart {
				words = append(words, capitalizeFirst(input[wordStart:i]))
			}
			wordStart = i + 1
		}
	}
	if wordStart < len(input) {
		words = append(words, capitalizeFirst(input[wordStart:]))
	}
	return joinStrings(words, " ")
}

func changeCase(input string) string {
	result := ""
	for _, ch := range input {
		if ch >= 'a' && ch <= 'z' {
			result += string(ch - 'a' + 'A')
		} else if ch >= 'A' && ch <= 'Z' {
			result += string(ch - 'A' + 'a')
		} else {
			result += string(ch)
		}
	}
	return result
}

func joinStrings(strings []string, sep string) string {
	if len(strings) == 0 {
		return ""
	}
	if len(strings) == 1 {
		return strings[0]
	}
	result := strings[0]
	for i := 1; i < len(strings); i++ {
		result += sep + strings[i]
	}
	return result
}
*/

/*
func CheckNumber(arg string) bool {
	for i := 0; i < len(arg); i++ {
		if isDigit(arg[i]) {
			return true
		}
	}
	return false
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
*/

/*
func ConcatAlternate(slice1, slice2 []int) []int {
	result := make([]int, 0)
	len1 := len(slice1)
	len2 := len(slice2)
	maxLen := len1
	if len2 > len1 {
		maxLen = len2
	}

	for i := 0; i < maxLen; i++ {
		if i < len1 {
			result = append(result, slice1[i])
		}
		if i < len2 {
			result = append(result, slice2[i])
		}
	}

	return result
}
*/

/*
func ConcatSlice(slice1, slice2 []int) []int {
	result := append(slice1, slice2...)
	return result
}
*/

/*
func CountAlpha(s string) int {
	count := 0
	for _, ch := range s {
		if isASCIIAlpha(ch) {
			count++
		}
	}
	return count
}

func isASCIIAlpha(ch rune) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z')
}
*/

/*
func CountChar(str string, c rune) int {
	count := 0
	for _, ch := range str {
		if ch == c {
			count++
		}
	}
	return count
}
*/

/*
func DigitLen(n, base int) int {
	if base < 2 || base > 36 {
		return -1
	}

	count := 0
	absN := n
	if n < 0 {
		absN = -n
	}

	for absN > 0 {
		absN /= base
		count++
	}

	return count
}
/*
func FifthAndSkip(str string) string {
	if len(str) < 5 {
		return "Invalid Input\n"
	}

	var builder []byte
	for i := 0; i < len(str); i += 5 {
		end := i + 5
		if end > len(str) {
			end = len(str)
		}

		substr := str[i:end]
		if len(substr) == 6 {
			substr = substr[:5]
		}

		builder = append(builder, substr...)
		builder = append(builder, ' ')
	}

	return string(builder[:len(builder)-1])
}
*/

/*
func FindPrevPrime(nb int) int {
	for i := nb; i >= 2; i-- {
		if isPrime(i) {
			return i
		}
	}
	return 0
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
*/

/*
func HalfSlice(slice []int) []int {
	length := len(slice)
	halfLength := (length + 1) / 2
	return slice[:halfLength]
}
*/

/*
func HashCode(dec string) string {
	size := len(dec)
	hashed := ""

	for _, char := range dec {
		hash := (int(char) + size) % 127
		if hash < 32 || hash > 126 {
			hash += 33
		}
		hashed += string(rune(hash))
	}

	return hashed
}
*/

/*
func IsSorted(f func(a, b int) int, a []int) bool {
	n := len(a)
	if n <= 1 {
		return true
	}

	for i := 1; i < n; i++ {
		if f(a[i-1], a[i]) > 0 {
			return false
		}
	}

	return true
}

func f(a, b int) int {
	return a - b
}
*/

/*
func ItoaBase(value, base int) string {
	// Handle zero separately
	if value == 0 {
		return "0"
	}

	// Handle negative numbers
	sign := ""
	if value < 0 {
		sign = "-"
		value = -value
	}

	// Define the characters used for each digit in the specified base
	digits := "0123456789ABCDEF"

	// Convert the number to a string in the specified base
	result := ""
	for value > 0 {
		digit := value % base
		result = string(digits[digit]) + result
		value /= base
	}

	return sign + result
}
*/

/*
func MultOrSum(ints []int, init int) int {
	// Check if the slice is empty
	if len(ints) == 0 {
		return 0
	}

	// Accumulated value
	accumulated := init

	// Iterate over the slice
	for _, num := range ints {
		// Check if the number is odd
		if num%2 != 0 {
			accumulated *= num
		} else {
			accumulated += num
		}
	}

	return accumulated
}
*/

/*
func displayLastWord(str string) {
	words := make([]string, 0)
	currWord := ""

	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			if currWord != "" {
				words = append(words, currWord)
				currWord = ""
			}
		} else {
			currWord += string(str[i])
		}
	}

	if currWord != "" {
		words = append(words, currWord)
	}

	if len(words) > 0 {
		lastWord := words[len(words)-1]
		fmt.Println(lastWord)
	}
}
*/

/*
func displayUniqueCharacters(str1, str2 string) {
	charSet := make(map[rune]bool)

	for _, char := range str1 {
		charSet[char] = true
	}

	for _, char := range str2 {
		charSet[char] = true
	}

	result := ""
	for char := range charSet {
		result += string(char)
	}

	fmt.Println(result)
}
*/

/*
func alphamirror(str string) {
	result := ""
	for _, char := range str {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			diff := int32('a')
			if char >= 'A' && char <= 'Z' {
				diff = int32('A')
			}
			opposite := diff + ('z' - (char - diff))
			result += string(opposite)
		} else {
			result += string(char)
		}
	}
	fmt.Println(result)
}
*/

/*
func findPrimeFactors(n int) []int {
	factors := []int{}
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
}
*/

/*
func SaveAndMiss(arg string, num int) string {
	if num <= 0 {
		return arg
	}

	var result string
	for i := 0; i < len(arg); i += num * 2 {
		end := i + num
		if end > len(arg) {
			end = len(arg)
		}
		result += arg[i:end]
	}

	return result
}
*/

/*
func canRewrite(firstStr, secondStr string) bool {
	firstIndex := 0
	secondIndex := 0

	for firstIndex < len(firstStr) && secondIndex < len(secondStr) {
		if firstStr[firstIndex] == secondStr[secondIndex] {
			firstIndex++
		}
		secondIndex++
	}

	return firstIndex == len(firstStr)
}
*/

/*
func ThirdTimeIsACharm(str string) string {
	if len(str) == 0 || len(str) < 3 {
		return "\n"
	}

	result := ""
	for i := 2; i < len(str); i += 3 {
		result += string(str[i])
	}

	return result + "\n"
}
*/

/*
func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}

	seen := make(map[rune]int)
	for _, ch := range str1 {
		seen[ch]++
	}

	for _, ch := range str2 {
		seen[ch]++
	}

	count := 0
	for _, value := range seen {
		if value == 1 {
			count++
		}
	}

	return count
}
*/

/*
func PrintIf(str string) string {
	if len(str) > 3 {
		return "G\n"
	} else if str == "" {
		return "G\n"
	} else {
		return "Invalid Input\n"
	}
}
*/

/*
func PrintIfNot(str string) string {
	if len(str) < 3 {
		return "G"
	} else if str == "" {
		return "G"
	} else {
		return "Invalid Input\n"
	}
}
*/

/*
func RectPerimeter(w, h int) int {
	if w < 0 || h < 0 {
		return -1
	}
	return 2 * (w + h)
}
*/

/*
func ReduceInt(a []int, f func(int, int) int) {
	if len(a) == 0 {
		return
	}

	result := a[0]
	for i := 1; i < len(a); i++ {
		result = f(result, a[i])
	}

	fmt.Println(result)
}
*/
