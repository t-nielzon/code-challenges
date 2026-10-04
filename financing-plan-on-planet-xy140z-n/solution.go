package main

func finance(n int) int64 {
	x := int64(n)
	return x * (x + 1) * (x + 2) / 2
}