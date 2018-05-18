package main

import (
	"bufio"
	"os"
)

type Job struct {
	id   int
	line string
}
type Result struct {
	job          Job
	lastSyllable string
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)
var lastSyllables = make(map[int]string)

func getFilename() string {
	reader := bufio.NewReader(os.Stdin)
	print("Enter the filepath of your poem: ")
	text, err := reader.ReadString('\n')
	check(err, "I can't read your input. Your typing must be terrible!")
	cleanFilename := text[:len(text)-1]
	return cleanFilename
}

func main() {
	println("Welcome to Angelou, a Golang rhymescheme generator!")
	poem, err := os.Open(getFilename())
	check(err, "I can't seem to open your file. Are you sure the filepath is correct?")
	defer poem.Close()
	scanner := bufio.NewScanner(poem)
	check(scanner.Err(), "I can't seem to read your file. Are you sure it hasn't been corrupted?")
	go allocate(scanner)
	done := make(chan bool)
	go result(done)
	createWorkerPool(10)
	<-done
	saveScheme(generateScheme(lastSyllables))
	println("I have saved your rhyme scheme in ./angelou.txt. Enjoy!")
}
