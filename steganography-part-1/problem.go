/*
## Steganography Part 1

Concealing a message within RGB image pixels using Least Significant Bit (LSB) Insertion.

Your function takes a message string and builds a bitmask array that is applied to the pixels array,
returning the modified pixel values.

Notes:
- pixels inner arrays always have length 3 (Red, Green, Blue)
- Return nil if there aren't enough pixels to conceal the message
- Remaining pixels after the message should be left unmodified
- The last color of every 3rd pixel (B channel of pixels 2, 5, 8, etc.) should not be modified
- Each character is converted to 8-bit binary and applied MSB-first to color values
*/

package main

func EncodeSteganography(msg string, pixels [][]int) [][]int {
	// placeholder
}