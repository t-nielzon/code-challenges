// solution.go
package main

import (
	"sort"
	"strings"
)

func DetectBruteForce(logs []string) []string {
	failureCounts := make(map[string]int)
	suspiciousIPs := make(map[string]bool)

	for _, log := range logs {
		parts := strings.Fields(log)
		if len(parts) < 2 {
			continue
		}

		ip := parts[0]
		status := parts[1]

		if status == "LOGIN_FAIL" {
			failureCounts[ip]++
			if failureCounts[ip] >= 3 {
				suspiciousIPs[ip] = true
			}
		} else if status == "LOGIN_SUCCESS" {
			failureCounts[ip] = 0
		}
	}

	result := make([]string, 0, len(suspiciousIPs))
	for ip := range suspiciousIPs {
		result = append(result, ip)
	}
	sort.Strings(result)

	return result
}