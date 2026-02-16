package kata

import "sort"

func Nico(key, message string) string {
	k := len(key)

	// pad message to a multiple of key length
	for len(message)%k != 0 {
		message += " "
	}

	// build numeric key: rank each character by its sorted position
	type pair struct {
		ch  byte
		idx int
	}
	pairs := make([]pair, k)
	for i := 0; i < k; i++ {
		pairs[i] = pair{key[i], i}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].ch < pairs[j].ch
	})

	order := make([]int, k)
	for rank, p := range pairs {
		order[p.idx] = rank
	}

	result := make([]byte, len(message))
	for i := 0; i < len(message); i += k {
		for j := 0; j < k; j++ {
			result[i+order[j]] = message[i+j]
		}
	}

	return string(result)
}