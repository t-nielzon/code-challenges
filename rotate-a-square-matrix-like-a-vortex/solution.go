package kata

func Vortex(matrix [][]int) [][]int {
	n := len(matrix)
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
		copy(result[i], matrix[i])
	}

	numRings := n / 2
	for ring := 0; ring < numRings; ring++ {
		steps := ring + 1
		elements := extractRing(result, ring, n)
		l := len(elements)
		steps = steps % l
		rotated := append(elements[steps:], elements[:steps]...)
		placeRing(result, ring, n, rotated)
	}

	return result
}

func extractRing(m [][]int, ring, n int) []int {
	var elems []int
	last := n - 1 - ring

	// top row left to right
	for j := ring; j < last; j++ {
		elems = append(elems, m[ring][j])
	}
	// right column top to bottom
	for i := ring; i < last; i++ {
		elems = append(elems, m[i][last])
	}
	// bottom row right to left
	for j := last; j > ring; j-- {
		elems = append(elems, m[last][j])
	}
	// left column bottom to top
	for i := last; i > ring; i-- {
		elems = append(elems, m[i][ring])
	}

	return elems
}

func placeRing(m [][]int, ring, n int, elems []int) {
	idx := 0
	last := n - 1 - ring

	for j := ring; j < last; j++ {
		m[ring][j] = elems[idx]
		idx++
	}
	for i := ring; i < last; i++ {
		m[i][last] = elems[idx]
		idx++
	}
	for j := last; j > ring; j-- {
		m[last][j] = elems[idx]
		idx++
	}
	for i := last; i > ring; i-- {
		m[i][ring] = elems[idx]
		idx++
	}
}