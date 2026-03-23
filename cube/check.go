package cube

// IsSolved checks if the cube is in solved state
func (c *Cube) IsSolved() bool {
	// Check if each face has uniform color
	for f := 0; f < 6; f++ {
		first := c.Stickers[f*FaceSize]
		for i := 1; i < FaceSize; i++ {
			if c.Stickers[f*FaceSize+i] != first {
				return false
			}
		}
	}
	return true
}
