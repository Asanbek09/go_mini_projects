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
	for name, testCase := range tests{
		t.Run(name, func(t *testing.T) {
			got, err := loadBookworms(testCase.bookwormsFile)
			if err != nil && !testCase.wantErr {
				t.Fatalf("expected an error %s, got none", err.Error())
			}

			if err == nil && testCase.wantErr {
				t.Fatalf("expected no error, got one %s", err.Error())
			}

			if !equalBookworms(got, testCase.want) {
				t.Fatalf("different result: got %v, expected: %v", got, testCase.want)
			}
		})
	}
}

func equalBookworms(t *testing.T, bookworms, target []Bookworm) bool {
	t.Helper()

	if len(bookworms) != len(target) {
		return false
	}

	for i := range bookworms {
		if bookworms[i].Name != target[i].Name {
			return false
		}

		if !equalBooks(t, bookworms[i].Books, target[i].Books) {
			return false
		}
	}

	return true
}