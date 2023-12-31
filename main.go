package main

import (
	"fmt"
	"sort"
)

type Dictionary map[string]string

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}

func (d Dictionary) Get(word string) string {
	return d[word]
}

func (d Dictionary) Remove(word string) {
	delete(d, word)
}

func (d Dictionary) List() []string {
	var result []string
	for word, definition := range d {
		result = append(result, fmt.Sprintf("%s: %s", word, definition))
	}
	sort.Strings(result)
	return result
}

func main() {
	dictionary := make(Dictionary)
	dictionary.Add("go", "A programming language.")
	dictionary.Add("Chaimae", "Prénom.")

	wordToGet := "go"
	fmt.Printf("Definition of %s: %s\n", wordToGet, dictionary.Get(wordToGet))

	fmt.Println("\nDictionary entries:")
	entries := dictionary.List()
	for _, entry := range entries {
		fmt.Println(entry)
	}
}
