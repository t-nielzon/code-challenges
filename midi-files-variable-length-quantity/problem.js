// MIDI Files: Variable Length Quantity
// 
// In MIDI files, timestamps are stored as integers representing the number of "ticks" since the previous event.
// MIDI uses Variable Length Quantity (VLQ) encoding to save space:
// - The lower 7 bits of each byte contain the actual data
// - The highest bit indicates whether another byte follows (1 = more bytes, 0 = final byte)
//
// Task: Implement two functions:
// - encode(n): convert an integer to a MIDI VLQ byte sequence
// - decode(data): decode only the first complete MIDI VLQ from the input
//
// All test values satisfy: 0 <= n <= 0x0FFFFFFF

function encode(n) {
  
}

function decode(data) {
  
}