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

func listOtherBooksOnShelves(bookIndexToRemove int, myBooks []Book) []Book {
	otherBooksOnShelves := make([]Book, bookIndexToRemove, len(myBooks)-1)
	copy(otherBooksOnShelves, myBooks[:bookIndexToRemove])
	otherBooksOnShelves = append(otherBooksOnShelves, myBooks[bookIndexToRemove+1:]...)

	return otherBooksOnShelves
}

func recommendBooks(recommendations bookRecommendations, myBooks []Book) []Book {
	bc := make(bookCollection)

	myShelf := make(map[Book]bool)
	for _, myBook := range myBooks {
		myShelf[myBook] = true
	}

	for _, myBook := range myBooks {
		for recommendation := range recommendations[myBook] {
			if myShelf[recommendation] {
				continue
			}

			bc[recommendation] = struct{}{}
		}
	}

	recommendationsForABook := bookCollectionToListOfBooks(bc)

	return recommendationsForABook
}