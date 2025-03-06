package solver

import (
	"fmt"
	"image"
)

type Solver struct {
	maze *image.RGBA
	palette palette
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

func (s *Solver) findEntrance() (image.Point, error) {
	for row := s.maze.Bounds().Min.Y; row < s.maze.Bounds().Max.Y; row++ {
		for col := s.maze.Bounds().Min.X; col < s.maze.Bounds().Max.X; col++ {
			if s.maze.RGBAAt(col, row) == s.palette.entrance {
				return image.Point{X: col, Y: row}, nil
			}
		}
	}

	return image.Point{}, fmt.Errorf("entrance position not found")
}