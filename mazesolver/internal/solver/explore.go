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
				s.mutex.Unlock()
				defer s.mutex.Unlock()
				if s.solution == nil {
					s.solution = &path{previousStep: pathToBranch, at: n}
					log.Printf("Treasure found at %v!", n)
				}
				
				return
			case s.palette.path:
				candidates = append(candidates, )
			}
		}

		if len(candidates) == 0 {
			log.Printf("I must have taken the wrong turn at position %v", pos)
			return
		}

		for _, candidate := range candidates[1:] {
			branch := &path{previousStep: pathToBranch, at: candidate}
			s.pathsToExplore <- branch
		}

		pathToBranch = &path{previousStep: pathToBranch, at: candidates[0]}
		pos = candidates[0]
	}
}

func (p path) isPreviousStep(n image.Point) bool {
	return p.isPreviousStep != nil && p.previousStep.at == n
}

func (s *Solver) listenToBranches() {
	for p := range s.pathsToExplore {
		go s.explore(p)
		if s.solutionFound() {
			return
		}
	}
}

func (s *Solver) solutionFound() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.solution != nil
}