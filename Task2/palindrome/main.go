package main

import (
	"fmt"
	"regexp"
	"strings"
)

func palindromeChecker(s string) bool {
	re := regexp.MustCompile(`[^\w\s]`)
	s = re.ReplaceAllLiteralString(s, "")
	s = strings.ToLower(s)

	i := 0
	j := len(s) - 1

	for i <= j{
		if s[i] != s[j]{
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	s := "ababaa"
	palindrome := palindromeChecker(s)
	if palindrome{
		fmt.Printf("the string %s is palindrome.", s)
	}else{
		fmt.Printf("the string %s is not palindrome.", s)
	}

}