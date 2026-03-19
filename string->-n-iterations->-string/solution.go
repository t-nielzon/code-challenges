package kata

func JumbledString(s string, n int) string {
	if len(s) <= 1 || n == 0 {
		return s
	}

	// find the cycle length: the permutation eventually returns to the original
	l := len(s)
	perm := make([]int, l)
	for i := 0; i < l; i++ {
		if i < (l+1)/2 {
			perm[i] = i * 2
		} else {
			perm[i] = (i-(l+1)/2)*2 + 1
		}
	}

	// find cycle length by tracking when index 0 returns
	// (LCM of all cycle lengths, but detecting full identity is safer)
	cycleLen := 1
	temp := make([]int, l)
	copy(temp, perm)
	for {
		identity := true
		for i := 0; i < l; i++ {
			if temp[i] != i {
				identity = false
				break
			}
		}
		if identity {
			break
		}
		// apply perm once more: temp = temp ∘ perm means temp[i] = perm[temp[i]]
		next := make([]int, l)
		for i := 0; i < l; i++ {
			next[i] = perm[temp[i]]
		}
		temp = next
		cycleLen++
	}

	n = n % cycleLen
	if n == 0 {
		return s
	}

	// binary exponentiation of the permutation
	bs := []byte(s)
	result := make([]byte, l)

	// compute perm^n via repeated squaring
	pow := make([]int, l)
	for i := range pow {
		pow[i] = i
	}

	cur := make([]int, l)
	copy(cur, perm)

	for n > 0 {
		if n&1 == 1 {
			tmp := make([]int, l)
			for i := 0; i < l; i++ {
				tmp[i] = cur[pow[i]]
			}
			pow = tmp
		}
		tmp := make([]int, l)
		for i := 0; i < l; i++ {
			tmp[i] = cur[cur[i]]
		}
		cur = tmp
		n >>= 1
	}

	for i := 0; i < l; i++ {
		result[i] = bs[pow[i]]
	}

	return string(result)
}