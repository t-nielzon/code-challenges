package main

func epidemic(tm float64, n int, s0, i0, b, a float64) int {
	dt := tm / float64(n)
	
	s := s0
	i := i0
	r := 0.0
	
	maxI := i0
	
	for k := 0; k < n; k++ {
		sNew := s - dt*b*s*i
		iNew := i + dt*(b*s*i-a*i)
		rNew := r + dt*a*i
		
		s = sNew
		i = iNew
		r = rNew
		
		if i > maxI {
			maxI = i
		}
	}
	
	return int(maxI)
}