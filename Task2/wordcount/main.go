package main

import (
	"fmt"
	"regexp"
	"strings"
)

func wordFrequency(s string) map[string]int{
	// Remove punctuation marks(ignore punctuation marks) -  using regexp
	re := regexp.MustCompile(`[^\w\s]`)
	s = re.ReplaceAllLiteralString(s, "")

	// make it case insensetive
	s = strings.ToLower(s)

	splitted_words := strings.Fields(s)
	freq := make(map[string]int)

	for _, word := range splitted_words {
		freq[word]++
	}
	return freq
}

func main(){
	s := "Hello, Welcome to go."	
	wordFrequencyDict := wordFrequency(s)
	fmt.Println("The word frequency count is:", wordFrequencyDict)	
}

