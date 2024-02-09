package main

import (
	"flag"
	"fmt"
	"strings"
)

// nonmatching chars in a string - also noted that we would only return unique chars
// for each character we are testing,

func findNonMatchingChars(s string, char string) map[string]bool {
	nonMatching := make(map[string]bool)
	for _, c := range s {
		if string(c) != char {
			nonMatching[string(c)] = true
		}
	}
	return nonMatching
}

func parseArg(s string) {

}

func main() {
	var testString string
	var testChars string

	flag.StringVar(&testString, "string", "test", "string to check")
	flag.StringVar(&testChars, "chars", "", "chars to filter out, can be a comma or space separated list")
	flag.Parse()

	l := strings.Split(testChars, " ")
	l = strings.Split(testChars, ",")

	for _, i := range l {
		fmt.Printf("CHAR: %s\n", i)
	}
	fmt.Printf("test string: %s\n", testString)
	fmt.Printf("test chars: %s\n", testChars)
	//var all := make(map[string]map)

	nonMatching := findNonMatchingChars(testString, testChars)
	fmt.Printf("%+v\n", nonMatching)
	var z string
	for k, _ := range nonMatching {
		z = fmt.Sprintf(" %s", k)
	}
	fmt.Printf("%s\n", z)
}
