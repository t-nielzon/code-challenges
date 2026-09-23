package main

func Amidakuji(ladder []string) []int {
	if len(ladder) == 0 {
		return []int{}
	}

	width := len(ladder[0])
	positions := make([]int, width)
	for i := 0; i < width; i++ {
		positions[i] = i
	}

	for _, level := range ladder {
		for i := 0; i < len(level); i++ {
			if level[i] == '1' && i+1 < len(positions) {
				positions[i], positions[i+1] = positions[i+1], positions[i]
			}
		}
	}

	return positions
}