package main

import (
	"log"
	"regexp"
)

func alphaNumeric(str string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9./ ]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(str, "")
}

func findIndex(slice []string, target string) int {
	for i, b := range slice {
		if b == target {
			return i
		}
	}
	return -1
}
func niceASCII(ASCII int) int {
	cleanASCII := ASCII + 64
	if cleanASCII > 90 {
		cleanASCII += 7
	}
	return cleanASCII
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
