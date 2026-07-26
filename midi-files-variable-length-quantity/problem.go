package kata

/*
# Introduction

In MIDI files, timestamps are stored as integers representing the number of "ticks"
since the previous event in the file.

Because these timestamps are relative to one another, most values tend to be fairly
small. In order to save space, MIDI files use a special integer encoding format
called a Variable Length Quantity (VLQ).

A MIDI VLQ stores an integer using one or more bytes:
- The lower 7 bits of each byte contain the actual data
- The highest bit indicates whether another byte follows
  - 1 means another byte follows
  - 0 means this is the final byte

For example:
  0x7F -> 01111111

Since the highest bit is 0, this is the final byte. Removing the continuation bit
leaves: 1111111 = 127

Values larger than 127 require multiple bytes. For example:
  [0x81, 0x00] in binary becomes: 10000001 00000000

Removing the continuation bits gives: 0000001 0000000 = 128

# Task

Implement two functions:
- Encode: convert an integer into a MIDI VLQ byte sequence
- Decode: decode only the first complete MIDI VLQ in the input, ignoring bytes after

All test values satisfy: 0 <= n <= 0x0FFFFFFF

# Examples

Encode:
  Encode(0)    == [0x00]
  Encode(127)  == [0x7F]
  Encode(128)  == [0x81, 0x00]
  Encode(8192) == [0xC0, 0x00]

Decode:
  Decode([0x00])                   == 0
  Decode([0x7F])                   == 127
  Decode([0x81, 0x00])             == 128
  Decode([0xFF, 0xFF, 0xFF, 0x7F]) == 268435455
*/

func Encode(n int) []byte {
	// TODO
	return nil
}

func Decode(data []byte) int {
	// TODO
	return 0
}