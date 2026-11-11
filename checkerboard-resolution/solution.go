package main

func CountBlackSquares(width, height, resolution int64) int64 {
	if width == 0 || height == 0 {
		return 0
	}

	fullBlocksX := width / resolution
	fullBlocksY := height / resolution
	remainderW := width % resolution
	remainderH := height % resolution

	var total int64 = 0

	// Main region: fullBlocksX * fullBlocksY blocks of size resolution x resolution
	total += blackBlocks(fullBlocksX, fullBlocksY) * resolution * resolution

	// Right strip: remainderW * fullBlocksY blocks of size remainderW x resolution
	if remainderW > 0 {
		var blackCount int64
		if fullBlocksX%2 == 0 {
			blackCount = fullBlocksY / 2
		} else {
			blackCount = (fullBlocksY + 1) / 2
		}
		total += blackCount * remainderW * resolution
	}

	// Bottom strip: fullBlocksX * remainderH blocks of size resolution x remainderH
	if remainderH > 0 {
		var blackCount int64
		if fullBlocksY%2 == 0 {
			blackCount = fullBlocksX / 2
		} else {
			blackCount = (fullBlocksX + 1) / 2
		}
		total += blackCount * resolution * remainderH
	}

	// Corner: remainderW x remainderH block
	if remainderW > 0 && remainderH > 0 {
		if (fullBlocksX+fullBlocksY)%2 == 1 {
			total += remainderW * remainderH
		}
	}

	return total
}

// blackBlocks returns count of black squares in an m x n checkerboard
// where top-left is white and each unit is a single square
func blackBlocks(m, n int64) int64 {
	return ((n + 1) / 2) * (m / 2) + (n / 2) * ((m + 1) / 2)
}