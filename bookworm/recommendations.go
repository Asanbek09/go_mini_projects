package main

import "sort"

type bookRecommendations map[Book]bookCollection

type bookCollection map[Book]struct{}

func newCollection() bookCollection {
	return make(bookCollection)
}

func recommendOtherBooks(bookworms []Bookworm) []Bookworm {
	sh := make(bookRecommendations)

	for _, bookworm := range bookworms {
		for i, book := range bookworm.Books {
			otherBooksOnShelves := listOtherBooksOnShelves(i, bookworm.Books)
			registerBookRecommendations(sb, book, otherBooksOnShelves)
		}
	}

	recommendations := make([]Bookworm, len(bookworms))
	for i, bookworm := range bookworms {
		recommendations[i] = Bookworm{
			Name: bookworm.Name,
			Books: recommendBooks(sb, bookworm.Books),
		}
	}

	return recommendations
}

type set map[Book] struct {}

func (s set) Contains(b Book) bool {
	_, ok := s[b]
	return ok
}