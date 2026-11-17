package main

func EncodeSteganography(msg string, pixels [][]int) [][]int {
	// check if there are enough pixels: each group of 3 pixels provides 8 usable channels
	if len(pixels) < len(msg)*3 {
		return nil
	}

	// create a copy of pixels to modify
	result := make([][]int, len(pixels))
	for i := range pixels {
		result[i] = make([]int, 3)
		copy(result[i], pixels[i])
	}

	// flatten the result for easier indexing
	flat := make([]int, len(pixels)*3)
	idx := 0
	for i := range result {
		flat[idx] = result[i][0]
		flat[idx+1] = result[i][1]
		flat[idx+2] = result[i][2]
		idx += 3
	}

	// process each character in the message
	colorIdx := 0
	for _, ch := range msg {
		byte := byte(ch)

		// process each bit of the character (MSB to LSB)
		for bit := 7; bit >= 0; bit-- {
			// skip indices where (colorIdx % 3 == 2), which are B channels
			for colorIdx%3 == 2 {
				colorIdx++
				if colorIdx >= len(flat) {
					return nil
				}
			}

			if colorIdx >= len(flat) {
				return nil
			}

			// extract the bit value
			bitVal := (byte >> uint(bit)) & 1

			// apply LSB insertion
			if bitVal == 1 {
				flat[colorIdx] |= 1
			} else {
				flat[colorIdx] &= ^1
			}

			colorIdx++
		}
	}

	// unflatten back to result
	idx = 0
	for i := range result {
		result[i][0] = flat[idx]
		result[i][1] = flat[idx+1]
		result[i][2] = flat[idx+2]
		idx += 3
	}

	return result
}