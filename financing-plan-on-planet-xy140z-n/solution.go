package main

func Finance(n int) int64 {
	n64 := int64(n)
	return n64 * (n64 + 1) * (n64 + 2) / 2
}