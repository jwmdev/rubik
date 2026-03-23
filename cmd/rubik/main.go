package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/jwmdev/rubik/cube"
	"github.com/jwmdev/rubik/solver"
	"github.com/jwmdev/rubik/solver/reduction"
	"github.com/jwmdev/rubik/solver/yau"
)

func main() {
	methodName := flag.String("method", "reduction", "Solver method (reduction, yau)")
	scramble := flag.String("scramble", "", "Scramble moves (space separated)")
	flag.Parse()

	c := cube.NewCube()

	// Apply scramble if provided
	if *scramble != "" {
		moves := strings.Split(*scramble, " ")
		fmt.Printf("Scrambling with: %s\n", *scramble)
		c.Scramble(moves)
	} else {
		fmt.Println("No scramble provided. Solving a clean cube (or use -scramble \"R U F ...\")")
	}

	var s solver.Method

	switch strings.ToLower(*methodName) {
	case "reduction":
		s = reduction.NewSolver()
	case "yau":
		s = yau.NewSolver()
	default:
		fmt.Printf("Unknown method: %s. Available: reduction, yau\n", *methodName)
		os.Exit(1)
	}

	fmt.Printf("Solving using %s...\n", s.Name())
	moves, err := s.Solve(c)
	if err != nil {
		fmt.Printf("Error solving: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Solution found:")
	if len(moves) == 0 {
		fmt.Println("(Already solved)")
	} else {
		for _, m := range moves {
			fmt.Println(m)
		}
	}
}
