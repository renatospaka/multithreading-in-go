package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"sync"
)

var (
	matches   []string
	waitGrout = sync.WaitGroup{}
	lock      = sync.Mutex{}
)

func fileSearch(root string, filename string) {
	log.Println("Searching in", root)
	files, _ := ioutil.ReadDir(root)
	for _, file := range files {
		if strings.Contains(file.Name(), filename) {
			lock.Lock()
			matches = append(matches, filepath.Join(root, file.Name()))
			lock.Unlock()
		}
		if file.IsDir() {
			waitGrout.Add(1)
			go fileSearch(filepath.Join(root, file.Name()), filename)
		}
	}
	waitGrout.Done()
}

func main() {
	waitGrout.Add(1)
	go fileSearch("/home/renatospaka/dev/cursos", "README.md")
	waitGrout.Wait()

	for _, file := range matches {
		fmt.Println("Matched:", file)
	}
	fmt.Printf("Found: %d matches\n", len(matches))
}
