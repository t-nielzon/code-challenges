package kata

import "sort"

func HashRadSeq(nMax int, k int) int {
	type entry struct {
		n      int
		radical int
	}

	entries := make([]entry, 0, nMax)

	for n := 1; n <= nMax; n++ {
		rad := calculateRadical(n)
		entries = append(entries, entry{n, rad})
	}

	sort.Slice(entries, func(i, j int) bool {
		if entries[i].radical != entries[j].radical {
			return entries[i].radical < entries[j].radical
		}
		return entries[i].n < entries[j].n
	})

	return entries[k-1].n
}

func calculateRadical(n int) int {
	if n == 1 {
		return 1
	}

	radical := 1

	if n%2 == 0 {
		radical *= 2
		for n%2 == 0 {
			n /= 2
		}
	}

	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			radical *= i
			for n%i == 0 {
				n /= i
			}
		}
	}

	if n > 1 {
		radical *= n
	}

	return radical
}