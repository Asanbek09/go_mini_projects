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
			registerBookRecommendations(sh, book, otherBooksOnShelves)
		}
	}

	recommendations := make([]Bookworm, len(bookworms))
	for i, bookworm := range bookworms {
		recommendations[i] = Bookworm{
			Name: bookworm.Name,
			Books: recommendBooks(sh, bookworm.Books),
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

func registerBookRecommendations(recommendations bookRecommendations, reference Book, otherBooksOnShelves []Book) {
	for _, book := range otherBooksOnShelves {
		collection, ok := recommendations[reference]
		if !ok {
			collection = newCollection()
			recommendations[reference] = collection
		}

		collection[book] = struct{}{}
	}
}

func bookCollectionToListOfBooks(bc bookCollection) []Book {
	bookList := make([]Book, 0, len(bc))
	for book := range bc {
		bookList = append(bookList, book)
	}

	sort.Slice(bookList, func(i, j int) bool {
		if bookList[i].Author != bookList[j].Author {
			return bookList[i].Author < bookList[j].Author
		}
		return bookList[i].Title < bookList[j].Title
	})

	return bookList
}