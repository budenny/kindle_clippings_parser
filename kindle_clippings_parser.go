package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type notes map[string][]string

func buildNotes(lines []string) notes {
	notes := make(map[string][]string)

	const linesPerClipping = 5
	for i := 0; i < len(lines)-linesPerClipping; i += linesPerClipping {
		book := lines[i]
		note := lines[i+3]
		notes[book] = append(notes[book], note)
	}
	return notes
}

func printNotes(n notes) {
	for book := range n {
		fmt.Println(book)
		for _, note := range n[book] {
			fmt.Println("\t" + note)
		}
	}
}

func readLines(filePath string) []string {
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), "\r\n")

}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage:\n %s <path to My Clippings.txt>\n", path.Base(os.Args[0]))
		os.Exit(-1)
	}

	lines := readLines(os.Args[1])
	notes := buildNotes(lines)
	printNotes(notes)
}
