package kata

func Vortex(matrix [][]int) [][]int {
	n := len(matrix)
	if n == 0 {
		return [][]int{}
	}

	// Create a copy of the matrix
	result := make([][]int, n)
	for i := range matrix {
		result[i] = make([]int, n)
		copy(result[i], matrix[i])
	}

	// Process each ring
	numRings := n / 2
	for ring := 0; ring < numRings; ring++ {
		rotations := ring + 1 // outer ring rotates 1, next rotates 2, etc.
		rotateRing(result, ring, rotations)
	}

	return result
}

func rotateRing(matrix [][]int, ring int, rotations int) {
	n := len(matrix)
	size := n - 2*ring // size of the current ring's side

	if size <= 1 {
		return
	}

	// Extract the ring elements into a slice (going counter-clockwise starting from top-left)
	// Top row (left to right), right column (top to bottom), bottom row (right to left), left column (bottom to top)
	// But we want to rotate "to the left" which is counter-clockwise

	// Extract ring elements in order: top row, right col, bottom row (reversed), left col (reversed)
	elements := make([]int, 0)

	// Top row (left to right)
	for j := ring; j < n-ring; j++ {
		elements = append(elements, matrix[ring][j])
	}
	// Right column (top+1 to bottom)
	for i := ring + 1; i < n-ring; i++ {
		elements = append(elements, matrix[i][n-ring-1])
	}
	// Bottom row (right-1 to left)
	for j := n - ring - 2; j >= ring; j-- {
		elements = append(elements, matrix[n-ring-1][j])
	}
	// Left column (bottom-1 to top+1)
	for i := n - ring - 2; i > ring; i-- {
		elements = append(elements, matrix[i][ring])
	}

	perim := len(elements)
	if perim == 0 {
		return
	}

	// Rotate left by 'rotations' steps (rotating left means element at index i goes to index (i - rotations))
	rotations = rotations % perim
	rotated := make([]int, perim)
	for i := 0; i < perim; i++ {
		newIdx := (i - rotations + perim) % perim
		rotated[newIdx] = elements[i]
	}

	// Put elements back
	idx := 0
	// Top row (left to right)
	for j := ring; j < n-ring; j++ {
		matrix[ring][j] = rotated[idx]
		idx++
	}
	// Right column (top+1 to bottom)
	for i := ring + 1; i < n-ring; i++ {
		matrix[i][n-ring-1] = rotated[idx]
		idx++
	}
	// Bottom row (right-1 to left)
	for j := n - ring - 2; j >= ring; j-- {
		matrix[n-ring-1][j] = rotated[idx]
		idx++
	}
	// Left column (bottom-1 to top+1)
	for i := n - ring - 2; i > ring; i-- {
		matrix[i][ring] = rotated[idx]
		idx++
	}
}