package kata

func CountBlackSquares(width, height, resolution int64) int64 {
	if width == 0 || height == 0 {
		return 0
	}

	fullBlocksX := width / resolution
	fullBlocksY := height / resolution
	remainderX := width % resolution
	remainderY := height % resolution
	r := resolution

	var count int64

	// Full block region
	evenBlockRows := (fullBlocksY + 1) / 2
	oddBlockRows := fullBlocksY / 2
	evenBlockCols := (fullBlocksX + 1) / 2
	oddBlockCols := fullBlocksX / 2

	// Even block rows have black at odd columns
	count += evenBlockRows * oddBlockCols * r * r
	// Odd block rows have black at even columns
	count += oddBlockRows * evenBlockCols * r * r

	// Right partial strip (remainder_x > 0)
	if remainderX > 0 {
		if fullBlocksX%2 == 1 {
			count += ((fullBlocksY + 1) / 2) * remainderX * r
		} else {
			count += (fullBlocksY / 2) * remainderX * r
		}
	}

	// Bottom partial strip (remainder_y > 0)
	if remainderY > 0 {
		if fullBlocksY%2 == 1 {
			count += ((fullBlocksX + 1) / 2) * r * remainderY
		} else {
			count += (fullBlocksX / 2) * r * remainderY
		}
	}

	// Bottom-right corner (both remainders > 0)
	if remainderX > 0 && remainderY > 0 {
		if (fullBlocksY+fullBlocksX)%2 == 1 {
			count += remainderX * remainderY
		}
	}

	return count
}