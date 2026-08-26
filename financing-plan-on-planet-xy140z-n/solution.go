package main

func Finance(n int) int64 {
	nn := int64(n)
	return nn * (nn + 1) * (nn + 2) / 2
}