package main

import "strings"

func BirdMountain(mountain string) int {
	lines := strings.Split(strings.TrimSpace(mountain), "\n")
	if len(lines) == 0 {
		return 0
	}

	height := len(lines)
	width := 0
	for _, line := range lines {
		if len(line) > width {
			width = len(line)
		}
	}

	// Create grid, padding shorter lines with spaces
	grid := make([][]rune, height)
	for i, line := range lines {
		grid[i] = make([]rune, width)
		for j := 0; j < width; j++ {
			if j < len(line) {
				grid[i][j] = rune(line[j])
			} else {
				grid[i][j] = ' '
			}
		}
	}

	// Distance map: -2 for peaks, -1 for unvisited non-peaks, 0+ for visited distances
	dist := make([][]int, height)
	queue := make([]struct{ i, j int }, 0)

	for i := range dist {
		dist[i] = make([]int, width)
		for j := range dist[i] {
			if grid[i][j] == '^' {
				dist[i][j] = -2
			} else {
				dist[i][j] = -1
			}
		}
	}

	// Initialize BFS: add all boundary and space cells with distance 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if dist[i][j] == -1 && (i == 0 || i == height-1 || j == 0 || j == width-1 || grid[i][j] == ' ') {
				dist[i][j] = 0
				queue = append(queue, struct{ i, j int }{i, j})
			}
		}
	}

	// BFS to compute distances for all non-peak cells
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		i, j := curr.i, curr.j

		dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
		for _, d := range dirs {
			ni, nj := i+d[0], j+d[1]
			if ni >= 0 && ni < height && nj >= 0 && nj < width {
				if dist[ni][nj] == -1 { // unvisited non-peak
					dist[ni][nj] = dist[i][j] + 1
					queue = append(queue, struct{ i, j int }{ni, nj})
				}
			}
		}
	}

	// Find maximum height: for each peak, get min neighbor distance, then max across all peaks
	maxHeight := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if grid[i][j] == '^' {
				minNeighbor := width + height
				dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
				for _, d := range dirs {
					ni, nj := i+d[0], j+d[1]
					if ni >= 0 && ni < height && nj >= 0 && nj < width {
						if dist[ni][nj] >= 0 && dist[ni][nj] < minNeighbor {
							minNeighbor = dist[ni][nj]
						}
					}
				}
				if minNeighbor < width+height && minNeighbor > maxHeight {
					maxHeight = minNeighbor
				}
			}
		}
	}

	return maxHeight
}