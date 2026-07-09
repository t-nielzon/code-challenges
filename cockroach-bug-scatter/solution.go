package main

import "fmt"

func Cockroach(room []string) [10]int {
	height := len(room)

	var cockroaches []struct {
		x, y  int
		dir   rune
	}

	for y := 0; y < height; y++ {
		for x := 0; x < len(room[y]); x++ {
			ch := rune(room[y][x])
			if ch == 'U' || ch == 'D' || ch == 'L' || ch == 'R' {
				cockroaches = append(cockroaches, struct {
					x, y  int
					dir   rune
				}{x, y, ch})
			}
		}
	}

	result := [10]int{}

	for _, roach := range cockroaches {
		holeNum := simulate(room, roach.x, roach.y, roach.dir)
		if holeNum >= 0 {
			result[holeNum]++
		}
	}

	return result
}

func simulate(room []string, startX, startY int, startDir rune) int {
	x, y := startX, startY
	dir := startDir

	visited := make(map[string]bool)

	for {
		state := fmt.Sprintf("%d,%d,%c", x, y, dir)
		if visited[state] {
			return -1
		}
		visited[state] = true

		nx, ny := x, y
		switch dir {
		case 'U':
			ny--
		case 'D':
			ny++
		case 'L':
			nx--
		case 'R':
			nx++
		}

		ch := rune(room[ny][nx])

		if ch == '+' || ch == '|' || ch == '-' {
			dir = turnLeft(dir)
		} else if ch >= '0' && ch <= '9' {
			return int(ch - '0')
		} else {
			x, y = nx, ny
		}
	}
}

func turnLeft(dir rune) rune {
	switch dir {
	case 'U':
		return 'L'
	case 'L':
		return 'D'
	case 'D':
		return 'R'
	case 'R':
		return 'U'
	}
	return dir
}