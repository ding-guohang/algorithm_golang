package util

import "strconv"

func PrintIntArray(nums []int) {
	s := ""
	for _, word := range nums {
		s += ", " + strconv.FormatInt(int64(word), 10)
	}
	println(s)
}

func PrintStringArray(words []string) {
	s := ""
	for _, word := range words {
		s += ", " + word
	}
	println(s)
}
