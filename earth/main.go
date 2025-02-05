package main

import "fmt"

func main() {
	greeting := greet("en")
	fmt.Println(greeting)
}

type language string

var phrasebook = map[language]string{
	"el": "Χαίρετε Κόσμε", // greek
	"en": "Hello World", // english
	"fr": "Bonjour le monde", // french
	"he": "שלום עולם", // hebrew
	"ur": "ہیلو دنیا", // urdu
	"vi": "Xin chào Thế Giới", // vietnamese
}

func greet(l language) string {
	greeting, ok := phrasebook[l]
	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}

	return greeting
}