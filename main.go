package main

import (
	"bufio"
	"fmt"
	"log"
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
	fmt.Print("Enter the filepath of your poem: ")
	text, err := reader.ReadString('\n')
	check(err)
	cleanFilename := text[:len(text)-1]
	return cleanFilename
}

func main() {
	fmt.Println("Welcome to Angelou, a Golang rhymescheme generator!")
	poem, err := os.Open(getFilename())
	check(err)
	defer poem.Close()
	scanner := bufio.NewScanner(poem)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	go allocate(scanner)
	done := make(chan bool)
	go result(done)
	createWorkerPool(10)
	<-done
	saveScheme(generateScheme(lastSyllables))
	fmt.Println("I have saved your rhyme scheme in ./angelou.txt. Enjoy!")
}
