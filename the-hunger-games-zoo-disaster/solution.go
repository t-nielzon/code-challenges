package kata

import "strings"

func WhoEatsWho(zoo string) []string {
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

	animals := strings.Split(zoo, ",")
	result := []string{zoo}

	for {
		ate := false
		for i := 0; i < len(animals); i++ {
			predator := animals[i]
			prey, ok := diet[predator]
			if !ok {
				continue
			}
			// check left first
			if i > 0 && prey[animals[i-1]] {
				result = append(result, predator+" eats "+animals[i-1])
				animals = append(animals[:i-1], animals[i:]...)
				ate = true
				break
			}
			// then check right
			if i < len(animals)-1 && prey[animals[i+1]] {
				result = append(result, predator+" eats "+animals[i+1])
				animals = append(animals[:i+1], animals[i+2:]...)
				ate = true
				break
			}
		}
		if !ate {
			break
		}
	}

	result = append(result, strings.Join(animals, ","))
	return result
}