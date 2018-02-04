package main

import "fmt"

// checkPasswords accepts a slice of passwords to validate, a map of commonly used passwords, an ASCII only bool flag
// and a character limit integer. If any of the passwords to validate fails one of the NIST guideline parameters the
// password will be printed to screen along with the reason why it failed.
func checkPasswords( pws []string, cmnPws map[string]bool, asciiFlag bool, charLimit *int ) {
	for _, pw := range pws {
		if len(pw) < 8 {
			fmt.Printf("%v -> Error: Too Short\n", pw)
			continue
		}
		if len(pw) > *charLimit {
			fmt.Printf("%v -> Error: Too Long\n", pw)
			continue
		}
		if asciiFlag {
			if !isASCIICompliant(pw) {
				fmt.Printf("%v -> Error: Invalid Characters\n", pw)
				continue
			}
		}
		if cmnPws[pw] {
			fmt.Printf("%v -> Error: Too Common\n", pw)
			continue
		}
	}
}

// checks each character in the string to make sure it is ASCII compliant my make sure it is within [0,127] byte range.
// taken from the table located here https://ascii.cl/
func isASCIICompliant( s string ) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > 127 || s[i] < 0 {
			return false
		}
	}
	return true
}