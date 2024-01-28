package main

// пока не работает
import (
	"fmt"
	"unicode"
)

func main() {
	var n int
	var s string
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s)
		if len(s) < 4 {
			fmt.Println("-")
			continue
		}
		validNum(s)
	}
}

func validNum(str string) {
	j := 0
	temp := []rune(str)
	result := make([]string, 0, 32)
	for i := 0; i < len(temp) && i+3 < len(temp); i++ {
		if (i+4 < len(temp)) && !unicode.IsDigit(temp[i]) && unicode.IsDigit(temp[i+1]) && unicode.IsDigit(temp[i+2]) && !unicode.IsDigit(temp[i+3]) && !unicode.IsDigit(temp[i+4]) {
			if i+5 == len(str) {
				result = append(result, []string{str[i:]}...)
				i += 4
				j = i
				continue
			}
			result = append(result, []string{str[i : i+5]}...)
			i += 4
			j = i
			continue
		}
		if !unicode.IsDigit(temp[i]) && unicode.IsDigit(temp[i+1]) && !unicode.IsDigit(temp[i+2]) && !unicode.IsDigit(temp[i+3]) {
			result = append(result, []string{str[i : i+4]}...)
			i += 3
			j = i
			continue
		}
		fmt.Println("-")
		return

	}
	if j >= len(temp)-1 {
		for _, v := range result {
			fmt.Print(v, " ")
		}
		return
	}
	fmt.Println("-")
}
