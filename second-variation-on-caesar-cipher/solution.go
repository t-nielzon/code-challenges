package kata

import "strings"

func shiftChar(c byte, shift int) byte {
	if c >= 'a' && c <= 'z' {
		return byte((int(c-'a')+shift%26+26)%26) + 'a'
	}
	if c >= 'A' && c <= 'Z' {
		return byte((int(c-'A')+shift%26+26)%26) + 'A'
	}
	return c
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

func splitIntoFive(s string) []string {
	n := len(s)
	partLen := n / 5
	remainder := n % 5
	// parts 1..4 have length partLen + (1 if needed), part 5 gets the rest
	// we want part 5 as long as possible, so give extra chars to parts 1-4 only as needed
	// each of the first `remainder` parts gets partLen+1, the rest get partLen
	// but constraint: parts 1-4 are >= part 5, and parts 1-4 are equal among the longer ones
	// Actually: parts 1-4 longer, part 5 shorter or equal. Maximize part 5 length.
	// So: let size = ceil(n/5) for parts 1-4 if needed.
	// part5 = n - 4*size. We want size minimal such that part5 <= size.
	// size = ceil(n/5). Then part5 = n - 4*size.

	if n == 0 {
		return []string{}
	}

	size := (n + 4) / 5 // ceil(n/5)
	if partLen == 0 {
		size = partLen + 1
	} else {
		size = partLen
		if remainder > 0 {
			size = partLen + 1
		}
	}

	// part5len = n - 4*size
	part5len := n - 4*size
	if part5len < 0 {
		// reduce size
		size = n / 4
		part5len = n - 4*size
	}

	var result []string
	idx := 0
	for i := 0; i < 4; i++ {
		if idx < n {
			end := idx + size
			if end > n {
				end = n
			}
			result = append(result, s[idx:end])
			idx = end
		}
	}
	if idx < n {
		result = append(result, s[idx:])
	}
	// remove trailing empty string
	if len(result) > 0 && result[len(result)-1] == "" {
		result = result[:len(result)-1]
	}
	_ = part5len
	return result
}

func Encode(s string, shift int) []string {
	if len(s) == 0 {
		return []string{}
	}

	// build prefix
	first := toLower(s[0])
	second := shiftChar(first, shift)
	prefix := string([]byte{first, second})

	// encode message
	var buf strings.Builder
	for i := 0; i < len(s); i++ {
		buf.WriteByte(shiftChar(s[i], shift))
	}

	coded := prefix + buf.String()
	return splitIntoFive(coded)
}

func Decode(arr []string) string {
	if len(arr) == 0 {
		return ""
	}

	coded := strings.Join(arr, "")
	if len(coded) < 2 {
		return ""
	}

	first := coded[0]
	second := coded[1]

	// determine shift: second = shiftChar(first, shift)
	shift := int(second) - int(first)
	if shift < 0 {
		shift += 26
	}
	shift = shift % 26

	// decode the rest
	message := coded[2:]
	var buf strings.Builder
	for i := 0; i < len(message); i++ {
		buf.WriteByte(shiftChar(message[i], -shift))
	}

	return buf.String()
}