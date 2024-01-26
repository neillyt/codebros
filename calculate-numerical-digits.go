package main

import (
	"flag"
	"fmt"
)

// start a count at 0
// iterate over the chars in the string
// if the char is a digit, increment count
//	we can tell if it is a digit by creating a list of digits, 0-9
//  we iterate over the digits to see if we find a match, return true if so, else false
// return count

func isDigit(r rune) bool {
	digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, d := range digits {
		if string(r) == d {
			return true
		}
	}
	return false
}

func countDigits(s string) int32 {
	var count int32
	for _, c := range s {
		if isDigit(c) {
			count++
		}
	}
	return count
}

func main() {
	var testString string
	flag.StringVar(&testString, "string", "he12399sdfsd83282934s9", "string to test")
	flag.Parse()
	fmt.Printf("test string: %s\n", testString)
	fmt.Printf("this string has %d digits\n", countDigits(testString))
}
