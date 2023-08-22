package main

type Point struct {
	x int
	y int
}

// up down left & right
var dir = [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

func walk(maze []string, wall string, curr Point, end Point, seen [][]bool, path *[]Point) bool {
	// 1. Base Case
	// off the map
	if curr.x < 0 || curr.x >= len(maze[0]) || curr.y < 0 || curr.y >= len(maze) {
		return false
	}
	// on a wall
	if string(maze[curr.y][curr.x]) == wall {
		return false
	}
	// at the end
	if curr.x == end.x && curr.y == end.y {
		*path = append(*path, end)
		return true
	}
	// have seen this point
	if seen[curr.y][curr.x] {
		return false
	}

	// 3 recurse
	// pre
	seen[curr.y][curr.x] = true
	*path = append(*path, curr)

	// recurse
	for i := 0; i < len(dir); i++ {
		x, y := dir[i][0], dir[i][1]

		success := walk(maze, wall, Point{
			x: curr.x + x,
			y: curr.y + y,
		}, end, seen, path)

		if success {
			return true
		}
	}

	// post
	// popping from stack
	*path = (*path)[:len(*path)-1]

	return false
}

func MazeSolver(maze []string, wall string, start Point, end Point) []Point {

	var seen [][]bool

	// initialize slice, otherwise will get index out of range error
	for i := 0; i < len(maze); i++ {
		var row []bool
		for range maze[0] {
			row = append(row, false)
		}
		seen = append(seen, row)
	}

	var path []Point

	walk(maze, wall, start, end, seen, &path)

	return path
}
