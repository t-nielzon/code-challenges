package kata

func Epidemic(tm, n, s0, i0, b, a float64) int {
	dt := tm / n
	s := s0
	i := i0
	r := 0.0
	maxI := i

	for k := 0; k < int(n); k++ {
		newS := s - dt*b*s*i
		newI := i + dt*(b*s*i-a*i)
		newR := r + dt*a*i
		s = newS
		i = newI
		r = newR
		if i > maxI {
			maxI = i
		}
	}

	return int(maxI)
}