package main

func ProductFib(prod int64) [3]int64 {
	fib1, fib2 := int64(0), int64(1)
	
	for fib1*fib2 <= prod {
		if fib1*fib2 == prod {
			return [3]int64{fib1, fib2, 1}
		}
		fib1, fib2 = fib2, fib1+fib2
	}
	
	return [3]int64{fib1, fib2, 0}
}