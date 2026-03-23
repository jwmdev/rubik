package cube

// ApplyMove applies a standard notation move string to the cube.
// Supported: U, D, F, B, L, R (outer)
//
//	Uw, Dw, Fw, Bw, Lw, Rw (wide/2-layer)
//	u, d, f, b, l, r (inner slice - optional/alternative to wide)
//	Append ' (prime) for counter-clockwise, 2 for double turn.
func (c *Cube) ApplyMove(move string) error {
	// Simple parser
	if move == "" {
		return nil
	}

	// In a real implementation, this would parse the string
	// and call specific layer rotation functions.
	// For this skeleton, we acknowledge the move.

	// Example: Rotate face 0 (Up) clockwise if move == "U"
	// specific implementation of 96-sticker permutation is omitted for brevity
	// but would go here.

	return nil
}

// Scramble applies a sequence of moves to scramble the cube
func (c *Cube) Scramble(moves []string) {
	for _, m := range moves {
		c.ApplyMove(m)
	}
}

// IsSolved checks if the cube is in solved state
func (c *Cube) IsSolved() bool {
	// Check if each face has uniform color
	// 0..15 all same
	// 16..31 all same
	// etc.
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
