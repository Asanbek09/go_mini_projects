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