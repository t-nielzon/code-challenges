package kata

func FindMissingLetter(letters string) byte {
	for i := 0; i < len(letters)-1; i++ {
		if letters[i+1]-letters[i] != 1 {
			return letters[i] + 1
		}
	}
	return 0
}