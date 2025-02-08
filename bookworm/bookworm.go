package main

import "os"

type Bookworm struct {
	Name string
	Books string
}

func loadBookworms(filePath string) ([]Bookworm, error) {
	
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return nil, nil
}

