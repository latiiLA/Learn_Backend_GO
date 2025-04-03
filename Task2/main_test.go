package main

import (
	"reflect"
	"testing"
)

func TestWordFrequency(t *testing.T){
	tests := []struct{
		input string
		expected map[string]int
	}{
		{
			input: "Hello, world! Hello world.",
			expected: map[string]int{
				"hello": 2,
				"world" : 2,
			},
		},
		{
			input: "regexp itself does not handle case sensitivity.",
			expected: map[string]int{
				"regexp": 1,
				"itself" : 1,
				"does" : 1,
				"not" : 1,
				"handle" : 1,
				"case": 1,
				"sensitivity" : 1,
			},
		},
		{
			input: "Test, test, TEST!",
			expected: map[string]int{
				"test": 3,
			},
		},
		{
			input: "   Spaces   between    words   ",
			expected: map[string]int{
				"spaces": 1,
				"between": 1,
				"words": 1,
			},
		},
	}

	for _, test := range tests{		
		result := wordFrequency(test.input)

		if !reflect.DeepEqual(result, test.expected){
			t.Errorf("For input: %q, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}