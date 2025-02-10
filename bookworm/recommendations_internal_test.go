package main

import (
	"testing"
	"reflect"
)


var (
	theHelp = Book{Author: "Kathryn Stockett", Title: "The Help"}
	fairyTale = Book{Author: "Stephen King", Title: "Fairy Tale"}
)

func equals(t *testing.T, bookwormA, bookwormB []Bookworm) bool {
	t.Helper()

	for i := range bookwormA {
		if !reflect.DeepEqual(bookwormA[i], bookwormB[i]) {
			return false
		}
	}
	
	return true
}