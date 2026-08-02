// solution.go
package main

func ProductFib(prod int64) [3]interface{} {
	a, b := int64(0), int64(1)
	
	for a*b < prod {
		a, b = b, a+b
	}
	
	if a*b == prod {
		return [3]interface{}{a, b, true}
	}
	return [3]interface{}{a, b, false}
}