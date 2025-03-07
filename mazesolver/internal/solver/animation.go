package solver

func (s *Solver) countExplorablePixels() int {
	explorablePixels := 0
	for row :=0; row < s.maze.Bounds().Dy(); row++ {
		for col := 0; col < s.maze.Bounds().Dx(); col++ {
			if s.maze.RGBAAt(col, row) != s.palette.wall {
				explorablePixels++
			}
		}
	}
	return explorablePixels
}

func (s *Solver) registerExploredPixels() {
	const totalExpectedFrames = 30

	explorablePixels := s.countExplorablePixels()
	pixelsExplored := 0

	for {
		select {
		case <-s.quit:
			return
		case pos := <-s.exploredPixels:
			s.maze.Set(pos.X, pos.Y, s.palette.explored)
			pixelsExplored++
			if pixelsExplored%(explorablePixels/totalExpectedFrames) == 0 {
				s.drawCurrentFrameToGIF()
			}
		}
	}
}