package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMazeSolver(t *testing.T) {
	maze := []string{
		"xxxxxxxxxx x",
		"x        x x",
		"x        x x",
		"x xxxxxxxx x",
		"x          x",
		"x xxxxxxxxxx"}

	mazeResult := []Point{
		{x: 10, y: 0},
		{x: 10, y: 1},
		{x: 10, y: 2},
		{x: 10, y: 3},
		{x: 10, y: 4},
		{x: 9, y: 4},
		{x: 8, y: 4},
		{x: 7, y: 4},
		{x: 6, y: 4},
		{x: 5, y: 4},
		{x: 4, y: 4},
		{x: 3, y: 4},
		{x: 2, y: 4},
		{x: 1, y: 4},
		{x: 1, y: 5},
	}

	result := MazeSolver(maze, "x", Point{10, 0}, Point{x: 1, y: 5})

	if !reflect.DeepEqual(mazeResult, result) {
		t.Errorf("wrong")
	}

	// draw the maze
	fmt.Println("MAZE ")
	for row := range maze {
		for pos := 0; pos < len(maze[row]); pos++ {
			fmt.Printf(string(maze[row][pos]))
		}
		fmt.Println()
	}
	fmt.Printf("\r\n\r\n")

	fmt.Println("SOLUTION")
	for row := range maze {
		for pos := 0; pos < len(maze[row]); pos++ {
			mazePos := string(maze[row][pos])
			y := row
			x := pos

			for res := range result {
				if result[res].y == y && result[res].x == x {
					mazePos = "*"
				}
			}

			fmt.Printf(mazePos)

		}
		fmt.Println()
	}

}
