package main

import "fmt"

func ConnectingValues(arr [][]int, val int, coord []int) [][]int {
	result := [][]int{}
	
	if len(arr) == 0 || len(coord) < 2 {
		return result
	}
	
	row, col := coord[0], coord[1]
	
	if row < 0 || row >= len(arr) || col < 0 || col >= len(arr[row]) {
		return result
	}
	
	if arr[row][col] != val {
		return result
	}
	
	visited := make(map[string]bool)
	queue := [][]int{{row, col}}
	visited[fmt.Sprintf("%d,%d", row, col)] = true
	
	directions := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1},           {0, 1},
		{1, -1},  {1, 0},  {1, 1},
	}
	
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		
		r, c := current[0], current[1]
		result = append(result, []int{r, c})
		
		for _, dir := range directions {
			nr, nc := r+dir[0], c+dir[1]
			
			if nr >= 0 && nr < len(arr) && nc >= 0 && nc < len(arr[nr]) {
				key := fmt.Sprintf("%d,%d", nr, nc)
				
				if !visited[key] && arr[nr][nc] == val {
					visited[key] = true
					queue = append(queue, []int{nr, nc})
				}
			}
		}
	}
	
	return result
}