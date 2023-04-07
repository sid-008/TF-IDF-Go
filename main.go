package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

type Document struct {
	Content string
	Terms   map[string]int
}

func NewDoc(content string) *Document {
	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}
	content = re.ReplaceAllString(content, " ")
	doc := &Document{
		Content: content,
		Terms:   make(map[string]int),
	}
	terms := strings.Fields(doc.Content)
	for _, term := range terms {
		doc.Terms[term]++
	}
	return doc
}

func main() {
	doc1 := `Hello from something something, I do something something, there is something in the way hmmmm mmmmm something in the way`
	test := NewDoc(doc1)
	fmt.Println(test.Terms)

}
