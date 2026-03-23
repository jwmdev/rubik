package cube

import (
	"errors"
	"fmt"
)

// Move represents a face turn.
// It can be a face turn (1, -1, 2) or a wide turn.
type Move int

// Move constants
// 0-17: U, D, F, B, R, L layer moves with modifiers (Normal, Double, Prime)
const (
	MoveU Move = iota
	MoveU2
	MoveUPrime
	MoveD
	MoveD2
	MoveDPrime
	MoveF
	MoveF2
	MoveFPrime
	MoveB
	MoveB2
	MoveBPrime
	MoveR
	MoveR2
	MoveRPrime
	MoveL
	MoveL2
	MoveLPrime

	// Wide moves (2 layers)
	MoveUw
	MoveUw2
	MoveUwPrime
	MoveDw
	MoveDw2
	MoveDwPrime
	MoveFw
	MoveFw2
	MoveFwPrime
	MoveBw
	MoveBw2
	MoveBwPrime
	MoveRw
	MoveRw2
	MoveRwPrime
	MoveLw
	MoveLw2
	MoveLwPrime
)

func (m Move) String() string {
	if m < 18 {
		face := []string{"U", "D", "F", "B", "R", "L"}[m/3]
		mod := []string{"", "2", "'"}[m%3]
		return face + mod
	}
	// Wide moves
	m2 := m - 18
	face := []string{"Uw", "Dw", "Fw", "Bw", "Rw", "Lw"}[m2/3]
	mod := []string{"", "2", "'"}[m2%3]
	return face + mod
}

// Face returns the face index (0-5: U, D, F, B, R, L)
func (m Move) Face() int {
	if m < 18 {
		return int(m) / 3
	}
	return int(m-18) / 3
}

// Turns returns 1, 2, or -1 (prime)
func (m Move) Turns() int {
	var mod int
	if m < 18 {
		mod = int(m) % 3
	} else {
		mod = int(m-18) % 3
	}

	if mod == 0 {
		return 1
	} else if mod == 1 {
		return 2
	}
	return -1
}

// IsWide returns true if the move is a wide move
func (m Move) IsWide() bool {
	return m >= 18
}

// ParseMove parses a string (e.g. "R2", "Uw'") into a Move
func ParseMove(s string) (Move, error) {
	if len(s) == 0 {
		return 0, errors.New("empty move string")
	}

	var base Move
	var suffix string

	// 1. Identify Face
	switch s[0] {
	case 'U':
		base = MoveU
	case 'D':
		base = MoveD
	case 'F':
		base = MoveF
	case 'B':
		base = MoveB
	case 'R':
		base = MoveR
	case 'L':
		base = MoveL
	default:
		return 0, fmt.Errorf("invalid face: %s", string(s[0]))
	}

	// 2. Identify if Wide (w)
	if len(s) > 1 && s[1] == 'w' {
		base += 18 // Jump to wide section
		if len(s) > 2 {
			suffix = s[2:]
		}
	} else {
		if len(s) > 1 {
			suffix = s[1:]
		}
	}

	// 3. Identify Modifier (2 or ')
	turns := 1 // Default
	if suffix == "2" {
		turns = 2
	} else if suffix == "'" {
		turns = -1
	} else if suffix != "" {
		return 0, fmt.Errorf("invalid modifier: %s", suffix)
	}

	// 4. Adjust base for modifier
	// Order in constant block: Face, Face2, Face'
	if turns == 2 {
		base += 1
	} else if turns == -1 {
		base += 2
	}

	return base, nil
}

// ApplyMove applies a move to the cube.
// Note: This modifies the cube in place.
func (c *Cube) ApplyMove(m Move) {
	// TODO: Implement actual permutation logic for 96 stickers here.
	// For now, this is a placeholder.
}

// Scramble applies a sequence of moves defined by strings
func (c *Cube) Scramble(moves []string) {
	for _, s := range moves {
		if s == "" {
			continue
		}
		if m, err := ParseMove(s); err == nil {
			c.ApplyMove(m)
		} else {
			fmt.Printf("Warning: Invalid scramble move '%s'\n", s)
		}
	}
}
