/*
# Introduction

In MIDI files, timestamps are stored as integers representing the number of "ticks" since the previous event in the file.

Because these timestamps are relative to one another, most values tend to be fairly small. In order to save space, MIDI files use a special integer encoding format called a **Variable Length Quantity (VLQ)**.

A MIDI VLQ stores an integer using one or more bytes:

- The lower 7 bits of each byte contain the actual data
- The highest bit indicates whether another byte follows
    - `1` means another byte follows
    - `0` means this is the final byte

# Task

Implement two functions:

  encode(n)
  decode(data)

- encode should convert an integer into a MIDI VLQ byte sequence
- decode should decode only the first complete MIDI VLQ contained in the input.
  Any additional bytes after the terminating VLQ byte should be ignored.

All test values will satisfy: 0 <= n <= 0x0FFFFFFF
*/

package kata

func Encode(n int) []byte {
}

func Decode(data []byte) int {
}