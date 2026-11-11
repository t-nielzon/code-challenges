package main

func FindMissingLetter(letters string) rune {
	for i := 0; i < len(letters)-1; i++ {
		if letters[i+1]-letters[i] != 1 {
			return rune(letters[i] + 1)
		}
	}
	return 0
}