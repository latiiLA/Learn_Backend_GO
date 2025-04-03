package main

import (
	"reflect"
	"testing"
)

func TestPalindrome(t *testing.T){ 
	tests := []struct{
		input string
		expected bool
	}{
		{
			input: "222",
			expected: true,
		},
		{
			input: "abcdcba",
			expected: true,
		},
		{
			input: "ccccfg",
			expected: false,
		},
		{
			input: "ABba",
			expected: true,
		},
	}

	for _, test := range tests{
		result := palindromeChecker(test.input)
		
		if !reflect.DeepEqual(result, test.expected){
			t.Errorf("For input %q, for expected %v, but got %v", test.input, test.expected, result)
		}
	}
}