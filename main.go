package main

import "fmt"

func main() {
	var counter int
	var input string
	fmt.Println("Enter a string")
	fmt.Scanln(&input)
	fmt.Println("Enter number of times you want to shift")
	fmt.Scanln(&counter)
	for i := 0; i < counter; i++ {
		input = changePosition(input)
	}
	fmt.Println(input)

}

func changePosition(str string) string {
	sp, fp := 0, 0
	for i, s := range str {
		val := string(s)
		if isVowel(val) {
			sp, fp = i, i+1
			break
		}
	}

	for fp < len(str) {
		fpVal := string(str[fp])
		if isVowel(fpVal) {
			spVal := rune(str[sp])
			str = replaceAtIndex(str, rune(str[fp]), sp)
			str = replaceAtIndex(str, spVal, fp)
			sp = fp
		}
		fp++

	}
	return str
}

func isVowel(val string) bool {
	return val == "a" || val == "e" || val == "i" || val == "o" || val == "u"
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
