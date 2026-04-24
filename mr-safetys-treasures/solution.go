package kata

import "strings"

func Unlock(s string) string {
	keys := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var b strings.Builder
	for _, r := range strings.ToLower(s) {
		for i, k := range keys {
			if strings.ContainsRune(k, r) {
				b.WriteByte(byte('2' + i))
				break
			}
		}
	}
	return b.String()
}