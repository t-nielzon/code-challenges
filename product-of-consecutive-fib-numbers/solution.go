package main

func ProductFib(prod int64) [3]interface{} {
	fib1, fib2 := int64(0), int64(1)
	
	for fib1*fib2 < prod {
		fib1, fib2 = fib2, fib1+fib2
	}
	
	if fib1*fib2 == prod {
		return [3]interface{}{fib1, fib2, true}
	}
	return [3]interface{}{fib1, fib2, false}
}