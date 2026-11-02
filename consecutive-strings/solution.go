package main

func LongestConsec(strarr []string, k int) string {
	n := len(strarr)
	
	if n == 0 || k <= 0 || k > n {
		return ""
	}
	
	var longestStr string
	maxLen := 0
	
	for i := 0; i <= n-k; i++ {
		var current string
		
		for j := 0; j < k; j++ {
			current += strarr[i+j]
		}
		
		if len(current) > maxLen {
			maxLen = len(current)
			longestStr = current
		}
	}
	
	return longestStr
}