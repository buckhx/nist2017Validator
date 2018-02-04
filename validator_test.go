package main

import "testing"

var asciiStrings = []string {
	"abcdefg",
	"123545",
	" #&$^@)!)#*",
	}
var unicodeStrings = []string {
	"ΛΡΦΫβ", // Greek characters
	"ดฒญง", // Thai characters
	"𝛀𝛂𝛉𝚷", // Mathematical symbols
}


func TestIsASCIICompliant( t *testing.T ) {
	// check for errors in ASCII compliant strings
	for _, str := range asciiStrings {
		if !isASCIICompliant(str) {
			t.Errorf("failed on test case: %v", str)
		}
	}

	// check for errors rejecting non-ASCII compliant unicode strings
	for _, str := range unicodeStrings {
		if isASCIICompliant(str) {
			t.Errorf("failed on test case: %v", str)
		}
	}

}

