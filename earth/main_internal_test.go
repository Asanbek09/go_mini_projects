package main

import "testing"

func ExampleMain(){
	main()
}

func TestGreet_English(t *testing.T) {
	lang := language("en")
	want := "Hello World!!!!"

	got := greet(lang)

	if got != want {
		t.Errorf("expected: %q, got: %q", want, got)
	}
}

func TestGreet_French(t *testing.T) {
	lang := language("fr")
	want := "Bonjour le mon"

	got := greet(lang)

	if got != want {
		t.Errorf("expected: %q, got: %q", want, got)
	}
}

func TestGreet_Akkadian(t *testing.T) {
	lang := language("akk")
	want := ""

	got := greet(lang)

	if got != want {
		t.Errorf("expected: %q, got: %q", want, got)
	}
}