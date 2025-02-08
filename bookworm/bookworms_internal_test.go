package main

import "testing"

var (
	handmaidsTale = Book{Author: "Margaret Atwwod", Title: "The Handmaid's Tale"}
	oryxAndCrake = Book{Author: "Margaret Atwood", Title: "Oryx and Crake"}
	theBellJar = Book{Author: "Sylvia Plath", Title: "The Bell Jar"}
	janeEyre = Book{Author: "Charlotte Bronte", Title: "Jane Eyre"}
)

func TestLoadBookworms_Success(t *testing.T) {
	type testCase struct {
		bookwormsFile string
		want []Bookworm
		wantErr bool
	}

	tests := map[string]testCase{
		"file exists": {
			bookwormsFile: "data.json",
			want: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			wantErr: false,
		},
		"file doesn't exist": {
			bookwormsFile: "database/no_file_here.json",
			want: nil,
			wantErr: true,
		},
		"invalid JSON": {
			bookwormsFile: "database/invalid.json",
			want: nil,
			wantErr: true,
		},
	}
}