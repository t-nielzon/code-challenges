/*
Bananas

Given a string of letters a, b, n how many different ways can you make the word
"banana" by crossing out various letters and then reading left-to-right?

(Use - to indicate a crossed-out letter)
*/
package kata

func Bananas(s string) []string {
	target := "banana"
	var results []string
	
	var backtrack func(sIdx, tIdx int, current []byte)
	backtrack = func(sIdx, tIdx int, current []byte) {
		// If we've matched the entire target, fill remaining with dashes
		if tIdx == len(target) {
			result := make([]byte, len(s))
			copy(result, current)
			for i := sIdx; i < len(s); i++ {
				result[i] = '-'
			}
			results = append(results, string(result))
			return
		}
		
		// If we've exhausted the input string, return
		if sIdx == len(s) {
			return
		}
		
		// Option 1: Cross out current character (use dash)
		current[sIdx] = '-'
		backtrack(sIdx+1, tIdx, current)
		
		// Option 2: If current char matches target char, use it
		if s[sIdx] == target[tIdx] {
			current[sIdx] = s[sIdx]
			backtrack(sIdx+1, tIdx+1, current)
		}
	}
	
	current := make([]byte, len(s))
	backtrack(0, 0, current)
	
	return results
}