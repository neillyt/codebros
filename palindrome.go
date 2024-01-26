package main

import (
	"flag"
	"fmt"
)

// split the string in half and start moving in both directions,
// comparing characters until you find 2 that are not equal.
//
// mirrorarorrim
// mirror rorrim
//       ^
// if mirro[r] != [r]orrim then false
// if mirr[o]r != r[o]rrim then false
// if mir[r]or != ro[r]rim then false
// if mi[r]ror != ror[r]im then false
// if m[i]rror != rorr[i]m then false
// if [m]irror != rorri[m] then false
// etc.

// OR
// do it the easy way. reverse the string and see if it is equal to the string.
// lol wtf.

func len(s string) int {
	var count int
	for range s {
		count++
	}
	return count
}

func isPalindromeHard(s string) bool {
	middle := len(s) / 2
	remainder := len(s) % 2
	beg := middle
	end := middle

	if remainder == 0 {
		beg--
	}

	for ; beg > -1; beg-- {
		if s[beg] != s[end] {
			return false
		}
		end++
	}
	return true
}

func reverse(s string) string {
	var newString string
	for _, v := range s {
		newString = string(v) + newString
	}
	return newString
}

func isPalindromeEasy(s string) bool {
	return s == reverse(s)
}

func main() {
	var testString string
	flag.StringVar(&testString, "string", "hellolleh", "string to test if palindrome")
	flag.Parse()
	fmt.Printf("test string: %s\n", testString)
	fmt.Printf("is palindrome hard: %t\n", isPalindromeHard(testString))
	fmt.Printf("is paldindrome easy: %t\n", isPalindromeEasy(testString))
	fmt.Printf("is palindrome hard: %t\n", isPalindromeHard(testString))
	fmt.Printf("is paldindrome easy: %t\n", isPalindromeEasy(testString))
}
