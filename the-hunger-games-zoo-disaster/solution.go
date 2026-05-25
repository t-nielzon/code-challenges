package kata

import "strings"

func ZooDisaster(zoo string) []string {
	diet := map[string]map[string]bool{
		"antelope": {"grass": true},
		"big-fish": {"little-fish": true},
		"bug":      {"leaves": true},
		"bear":     {"big-fish": true, "bug": true, "chicken": true, "cow": true, "leaves": true, "sheep": true},
		"chicken":  {"bug": true},
		"cow":      {"grass": true},
		"fox":      {"chicken": true, "sheep": true},
		"giraffe":  {"leaves": true},
		"lion":     {"antelope": true, "cow": true},
		"panda":    {"leaves": true},
		"sheep":    {"grass": true},
	}

	items := strings.Split(zoo, ",")
	result := []string{zoo}

	for {
		ate := false
		for i := 0; i < len(items); i++ {
			eater := items[i]
			menu, ok := diet[eater]
			if !ok {
				continue
			}
			if i > 0 && menu[items[i-1]] {
				result = append(result, eater+" eats "+items[i-1])
				items = append(items[:i-1], items[i:]...)
				ate = true
				break
			}
			if i < len(items)-1 && menu[items[i+1]] {
				result = append(result, eater+" eats "+items[i+1])
				items = append(items[:i+1], items[i+2:]...)
				ate = true
				break
			}
		}
		if !ate {
			break
		}
	}

	result = append(result, strings.Join(items, ","))
	return result
}