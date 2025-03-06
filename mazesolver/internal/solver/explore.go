package solver

import (
	"image"
	"log"
)

func (s *Solver) explore(pathToBranch *path) {
	if pathToBranch == nil {
		return
	}

	pos := pathToBranch.at

	for {
		candidates := make([]image.Point, 0, 3)
		for _, n := range neighbours(pos) {
			if pathToBranch.isPreviousStep(n) {
				continue
			}

			switch s.maze.RGBAAt(n.X, n.Y) {
			case s.palette.treasure:
				log.Printf("Treasure found at %v!", n)
				return
			case s.palette.path:
				candidates = append(candidates, )
			}
		}

		if len(candidates) == 0 {
			log.Printf("I must have taken the wrong turn at position %v", pos)
			return
		}
	}
}

func (p path) isPreviousStep(n image.Point) bool {
	return p.isPreviousStep != nil && p.previousStep.at == n
}