// problem.go
package kata

// Given a semi-inclusive interval I = [l, u) where l is included but u is not,
// an integer n > 0, and a function f: x -> f(x),
// return a list of n values: f(l), f(l + 1*d), ..., f(u - d)
// where d = (u - l) / n
//
// Each resulting value should be floored to two decimals using: floor(y * 100.0) / 100.0
//
// Examples:
// interp(x -> x, 0.0, 0.9, 3) -> [0.0, 0.3, 0.6]
// interp(x -> x, 0.0, 0.9, 4) -> [0.0, 0.22, 0.45, 0.67]
// interp(x -> x, 0.0, 1.0, 4) -> [0.0, 0.25, 0.5, 0.75]
// interp(x -> sin x, 0.0, 0.9, 3) -> [0.0, 0.29, 0.56]

func interp(f func(float64) float64, l, u float64, n int) []float64 {
	// TODO
	return nil
}