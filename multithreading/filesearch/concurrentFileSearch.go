package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

var (
	matches []string
)

func fileSearch(root string, filename string) {
	log.Println("Searching in", root)
	files, _ := ioutil.ReadDir(root)
	for _, file := range files {
		if strings.Contains(file.Name(), filename) {
			matches = append(matches, filepath.Join(root, file.Name()))
		}
		if file.IsDir() {
			fileSearch(filepath.Join(root, file.Name()), filename)
		}
	} 
}

func main() {
	fileSearch("/home/renatospaka/dev/cursos", "README.md")
	for _, file := range matches {
		fmt.Println("Matched:", file)
	}
	fmt.Printf("Found: %d matches\n", len(matches))
}
