package solver

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"errors"
)

func openMaze(imagePath string) (*image.RGBA, error) {
	f, err := os.Open(imagePath)
	if err != nil {
		return nil, fmt.Errorf("unable to open image %s: %w", imagePath, err)
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("unable to load input image from %s: %w", imagePath, err)
	}

	rgbaImage, ok := img.(*image.RGBA)
	if !ok {
		return nil, fmt.Errorf("expected RGBA image, got %T", img)
	}

	return rgbaImage, nil
}

func (s *Solver) SaveSolution(outputPath string) (err error) {
	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("unable to create output image file at %s", outputPath)
	}
	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			err = errors.Join(err, fmt.Errorf("unable to close file: %w", closeErr))
		}
	}()
	
	stepsFromTreasure := s.solution
	for stepsFromTreasure != nil {
		s.maze.Set(stepsFromTreasure.at.X, stepsFromTreasure.at.Y, s.palette.solution)
		stepsFromTreasure = stepsFromTreasure.previousStep
	}

	err = png.Encode(f, s.maze)
	if err != nil {
		return fmt.Errorf("unable to write output image at %s: %w", outputPath, err)
	}

	return nil
}