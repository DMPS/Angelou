package main

import (
	"os"
	"strings"
)

func getLastSyllable(line string) string {
	if line == "" {
		return line
	}
	return line[len(line)-3:]
}

func generateScheme(poem map[int]string) string {
	scheme := strings.Builder{}
	seenSyllables := []string{}
	for i := 0; i < len(lastSyllables); i++ {
		syllable := poem[i]
		if syllable == "" {
			scheme.WriteString("\n")
		} else {
			if ascii := findIndex(seenSyllables, syllable); ascii == -1 {
				seenSyllables = append(seenSyllables, syllable)
				letter := rune(niceASCII(len(seenSyllables)))
				scheme.WriteRune(letter)
				scheme.WriteString("\n")
			} else {
				scheme.WriteRune(rune(niceASCII(ascii)))
				scheme.WriteString("\n")
			}
		}
	}
	return scheme.String()
}

func saveScheme(scheme string) {
	schemeFile, createErr := os.Create("./angelou.txt")
	check(createErr, "I can't write to the file. Do you need to give me write access?")
	defer schemeFile.Close()
	_, writeErr := schemeFile.WriteString(scheme)
	check(writeErr, "I can't write to the file. Do you need to give me write access?")
}
