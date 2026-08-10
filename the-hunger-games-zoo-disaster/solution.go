package main

import "strings"

func ZooDis(zoo string) []string {
	eats := map[string][]string{
		"antelope":   {"grass"},
		"big-fish":   {"little-fish"},
		"bug":        {"leaves"},
		"bear":       {"big-fish", "bug", "chicken", "cow", "leaves", "sheep"},
		"chicken":    {"bug"},
		"cow":        {"grass"},
		"fox":        {"chicken", "sheep"},
		"giraffe":    {"leaves"},
		"lion":       {"antelope", "cow"},
		"panda":      {"leaves"},
		"sheep":      {"grass"},
	}

	animals := strings.Split(zoo, ",")
	result := []string{zoo}

	canEat := func(predator string, food string) bool {
		foods, ok := eats[predator]
		if !ok {
			return false
		}
		for _, f := range foods {
			if f == food {
				return true
			}
		}
		return false
	}

	for {
		eaten := false

		for i := 0; i < len(animals) && !eaten; i++ {
			if i > 0 && canEat(animals[i], animals[i-1]) {
				result = append(result, animals[i]+" eats "+animals[i-1])
				animals = append(animals[:i-1], animals[i:]...)
				eaten = true
			} else if i < len(animals)-1 && canEat(animals[i], animals[i+1]) {
				result = append(result, animals[i]+" eats "+animals[i+1])
				animals = append(animals[:i+1], animals[i+2:]...)
				eaten = true
			}
		}

		if !eaten {
			break
		}
	}

	result = append(result, strings.Join(animals, ","))
	return result
}