package kata

func WhichNote(keyPressCount int) string {
	notes := []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	// 88 keys span 7 full octaves plus 3 extra keys (A, A#, B)
	key := ((keyPressCount - 1) % 88)
	return notes[key%12]
}