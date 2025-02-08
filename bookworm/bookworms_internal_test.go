package main

import "testing"

func TestLoadBookworms_Success(t *testing.T) {
	type testCase struct {
		bookwormsFile string
		want []Bookworm
		wantErr bool
	}
}

var (
	handmaidsTale = Book{Author: "Margaret Atwwod", Title: "The Handmaid's Tale"}
	oryxAndCrake = Book{Author: "Margaret Atwood", Title: "Oryx and Crake"}
)