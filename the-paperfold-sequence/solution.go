package main

// PaperFold returns a function that generates successive values of the paper folding sequence
func PaperFold() func() int {
	sequence := []int{1}
	index := 0

	return func() int {
		// expand sequence as needed
		for len(sequence) <= index {
			newSeq := make([]int, len(sequence)*2+1)
			copy(newSeq, sequence)
			newSeq[len(sequence)] = 1

			// mirror with bit flip: reverse(NOT(previous))
			prevLen := len(sequence)
			for i := 0; i < prevLen; i++ {
				newSeq[prevLen+1+i] = 1 - sequence[prevLen-1-i]
			}

			sequence = newSeq
		}

		result := sequence[index]
		index++
		return result
	}
}