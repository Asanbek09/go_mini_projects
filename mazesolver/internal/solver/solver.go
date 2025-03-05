package solver

import (
	"fmt"
	"image"
)

type Solver struct {
	maze *image.RGBA
}

func (s *Solver) Solve() error {
	return nil
}

func New(imagePath string) (*Solver, error) {
	img, err := openMaze(imagePath)
	if err != nil {
		return nil, fmt.Errorf("cannot open maze image: %w", err)
	}

	return &Solver{
		maze: img,
	}, nil
}