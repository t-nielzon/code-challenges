package main

import (
	"fmt"
	"strings"
)

func ZooDis(zoo string) []string {
	eats := map[string][]string{
		"antelope": {"grass"},
		"big-fish": {"little-fish"},
		"bug": {"leaves"},
		"bear": {"big-fish", "bug", "chicken", "cow", "leaves", "sheep"},
		"chicken": {"bug"},
		"cow": {"grass"},
		"fox": {"chicken", "sheep"},
		"giraffe": {"leaves"},
		"lion": {"antelope", "cow"},
		"panda": {"leaves"},
		"sheep": {"grass"},
	}
	
	result := []string{zoo}
	animals := strings.Split(zoo, ",")
	
	for {
		ate := false
		
		for i := 0; i < len(animals); i++ {
			eatable, ok := eats[animals[i]]
			if !ok {
				continue
			}
			
			// Check left first
			if i > 0 {
				for _, food := range eatable {
					if animals[i-1] == food {
						victim := animals[i-1]
						animals = append(animals[:i-1], animals[i:]...)
						result = append(result, fmt.Sprintf("%s eats %s", animals[i-1], victim))
						ate = true
						break
					}
				}
			}
			
			if ate {
				break
			}
			
			// Check right
			if i < len(animals)-1 {
				for _, food := range eatable {
					if animals[i+1] == food {
						victim := animals[i+1]
						animals = append(animals[:i+1], animals[i+2:]...)
						result = append(result, fmt.Sprintf("%s eats %s", animals[i], victim))
						ate = true
						break
					}
				}
			}
			
			if ate {
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