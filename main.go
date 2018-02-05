package main

import (
	"fmt"
	"bufio"
	"os"
	"flag"
	"log"
	"github.com/pkg/errors"
)

func main() {
	pwsFlag := flag.String("common-pws", "", "common passwords dictionary path")
	maxCharFlag := flag.Int("max-char", 128, "maximum numbe of characters allowed")
	asciiOnly := flag.Bool("ascii-only", true, "only allows ascii compliant passwords")
	flag.Parse()

	if *pwsFlag == "" {
		log.Fatal(errors.New("you must define a path for your list of common passwords to check, using " +
			"the flag --common-pws=$PATH"))
	}
	// enforce a minimum maximum character limit of 64
	if *maxCharFlag < 64 {
		fmt.Print("Maximum character limit cannot be less then 64, setting to 64.\n")
		*maxCharFlag = 64
	}

	// open common passwords file from user supplied path
	file, err := os.Open(*pwsFlag)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var cmnPws = make( map[string]bool )

	// build map of commonly used passwords for validation from open file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cmnPws[scanner.Text()] = true
	}

	// check for any errors during scan
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}


	var pws []string
	// build slice of passwords to validate from STDIN
	scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		pws = append(pws, scanner.Text())
	}

	// check for any errors during scan
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	checkPasswords(pws, cmnPws, *asciiOnly, maxCharFlag)
}

