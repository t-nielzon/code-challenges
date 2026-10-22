package main

func CountBlackSquares(width, height, resolution uint64) uint64 {
	if width == 0 || height == 0 {
		return 0
	}
	
	fullCols := width / resolution
	fullRows := height / resolution
	partialColWidth := width % resolution
	partialRowHeight := height % resolution
	
	numEvenRows := (fullRows + 1) / 2
	numOddRows := fullRows / 2
	numEvenCols := (fullCols + 1) / 2
	numOddCols := fullCols / 2
	
	var blackPixels uint64 = 0
	
	// Even rows have black at odd columns
	blackPixels += numEvenRows * numOddCols * resolution * resolution
	if partialColWidth > 0 && fullCols%2 == 1 {
		blackPixels += numEvenRows * partialColWidth * resolution
	}
	
	// Odd rows have black at even columns
	blackPixels += numOddRows * numEvenCols * resolution * resolution
	if partialColWidth > 0 && fullCols%2 == 0 {
		blackPixels += numOddRows * partialColWidth * resolution
	}
	
	// Partial row (if it exists)
	if partialRowHeight > 0 {
		if fullRows%2 == 0 {
			// Partial row is even
			blackPixels += numOddCols * resolution * partialRowHeight
			if partialColWidth > 0 && fullCols%2 == 1 {
				blackPixels += partialColWidth * partialRowHeight
			}
		} else {
			// Partial row is odd
			blackPixels += numEvenCols * resolution * partialRowHeight
			if partialColWidth > 0 && fullCols%2 == 0 {
				blackPixels += partialColWidth * partialRowHeight
			}
		}
	}
	
	return blackPixels
}