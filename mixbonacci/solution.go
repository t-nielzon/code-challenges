package kata

func Mixbonacci(p []string, n int) []int {
	if n == 0 || len(p) == 0 {
		return []int{}
	}

	generators := map[string]func() func() int{
		"fib": func() func() int {
			a, b := 0, 1
			return func() int {
				v := a
				a, b = b, a+b
				return v
			}
		},
		"pad": func() func() int {
			a, b, c := 1, 0, 0
			return func() int {
				v := a
				a, b, c = b, c, a+b
				return v
			}
		},
		"jac": func() func() int {
			a, b := 0, 1
			return func() int {
				v := a
				a, b = b, a+2*b
				return v
			}
		},
		"pel": func() func() int {
			a, b := 0, 1
			return func() int {
				v := a
				a, b = b, 2*b+a
				return v
			}
		},
		"tri": func() func() int {
			a, b, c := 0, 0, 1
			return func() int {
				v := a
				a, b, c = b, c, a+b+c
				return v
			}
		},
		"tet": func() func() int {
			a, b, c, d := 0, 0, 0, 1
			return func() int {
				v := a
				a, b, c, d = b, c, d, a+b+c+d
				return v
			}
		},
	}

	active := make(map[string]func() int)
	for _, key := range p {
		if _, ok := active[key]; !ok {
			active[key] = generators[key]()
		}
	}

	// Need per-position generators since same key can appear multiple times
	// Each position in pattern with same key shares the same generator
	// Actually, looking at the example: 'fib','fib','pel' shares the fib generator
	// So the map approach above is correct.

	result := make([]int, n)
	for i := 0; i < n; i++ {
		key := p[i%len(p)]
		result[i] = active[key]()
	}
	return result
}