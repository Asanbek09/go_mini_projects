package main


type testCase struct {
	bookwormsFile string
	want []Bookworm
	wantErr bool
}