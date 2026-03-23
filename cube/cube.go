package cube

import (
	"fmt"
	"strings"
)

type Color int

const (
	White Color = iota
	Yellow
	Green
	Blue
	Red
	Orange
)

// Face names
const (
	FaceU = "Up"
	FaceD = "Down"
	FaceF = "Front"
	FaceB = "Back"
	FaceL = "Left"
	FaceR = "Right"
)

// A 4x4x4 cube has 6 faces * 16 stickers each = 96 stickers total
// CubeSize is the dimension of the cube (4x4x4)
const CubeSize = 4
const FaceSize = CubeSize * CubeSize
const TotalStickers = 6 * FaceSize

type Cube struct {
	Stickers [TotalStickers]Color
}

func NewCube() *Cube {
	c := &Cube{}
	// Initialize solved state
	// Top (White) - 0..15
	// Bottom (Yellow) - 16..31
	// Front (Green) - 32..47
	// Back (Blue) - 48..63
	// Left (Orange) - 64..79
	// Right (Red) - 80..95

	for i := 0; i < FaceSize; i++ {
		c.Stickers[i] = White
		c.Stickers[i+FaceSize] = Yellow
		c.Stickers[i+2*FaceSize] = Green
		c.Stickers[i+3*FaceSize] = Blue
		c.Stickers[i+4*FaceSize] = Orange
		c.Stickers[i+5*FaceSize] = Red
	}
	return c
}

func (c *Cube) Clone() *Cube {
	newC := &Cube{}
	copy(newC.Stickers[:], c.Stickers[:])
	return newC
}

func (c *Cube) String() string {
	// Basic visualization or string representation
	var sb strings.Builder
	// Just dump raw values for now
	for i, s := range c.Stickers {
		sb.WriteString(fmt.Sprintf("%d", s))
		if (i+1)%FaceSize == 0 {
			sb.WriteString("\n")
		} else if (i+1)%4 == 0 {
			sb.WriteString(" ")
		}
	}
	return sb.String()
}
