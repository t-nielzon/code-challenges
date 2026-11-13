package main

func LongestConsec(strarr []string, k int) string {
	n := len(strarr)
	
	if n == 0 || k > n || k <= 0 {
		return ""
	}
	
	var longestStr string
	maxLen := 0
	
	for i := 0; i <= n-k; i++ {
		var concat string
		for j := i; j < i+k; j++ {
			concat += strarr[j]
		}
		
		if len(concat) > maxLen {
			maxLen = len(concat)
			longestStr = concat
		}
	}
	
	return longestStr
}