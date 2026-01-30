package kata

import "sort"

func FindTheKey(messages, secrets []string) string {
	pairs := make(map[string]bool)

	for i := 0; i < len(messages); i++ {
		msg := messages[i]
		sec := secrets[i]

		for j := 0; j < len(msg); j++ {
			m := msg[j]
			s := sec[j]

			if m != s && m != ' ' && s != ' ' {
				var pair string
				if m < s {
					pair = string(m) + string(s)
				} else {
					pair = string(s) + string(m)
				}
				pairs[pair] = true
			}
		}
	}

	pairList := make([]string, 0, len(pairs))
	for p := range pairs {
		pairList = append(pairList, p)
	}

	sort.Strings(pairList)

	result := ""
	for _, p := range pairList {
		result += p
	}

	return result
}