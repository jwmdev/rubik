package solver

import (
	"github.com/jwmdev/rubik/cube"
)

// Method is the interface for different solving methods
type Method interface {
	Solve(c *cube.Cube) ([]string, error)
	Name() string
}

// Result contains solving details
type Result struct {
	Moves []string
	Error error
}
