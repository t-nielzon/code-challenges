package main

import "fmt"

func IntToIP(num uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d",
		(num >> 24) & 0xFF,
		(num >> 16) & 0xFF,
		(num >> 8) & 0xFF,
		num & 0xFF)
}