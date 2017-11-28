package main

import (
	    "unicode"
	    "strings"
)

func RemoveEven(arr []int) (result []int) {
	for i, a := range arr {
		if a % 2 == 1 {
			result = append(result, arr[i])
		}
	}
	return
}

func PowerGenerator(a int) (func() int) {
	cur := 1
	return func() int {
		cur = cur * a
		return cur
	}
}

func DifferentWordsCount(str string) int {
	dict := make(map[string]bool)
	start := -1
	for i, c := range str {
		if !unicode.IsLetter(c) {
			if i - start > 1 {
				dict[strings.ToLower(string(str[start + 1:i]))] = true
			}
			start = i
		}

	}
	return len(dict)
}