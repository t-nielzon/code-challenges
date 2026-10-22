// Mirrored Exponential Chunks
// Using a single function provided with an array, create a new one in the following ways:
// - Separate array elements into chunks. Each chunk will contain 2^abs(n) elements where n is the "distance" to the median chunk.
// - The median chunk(s) will have one single element, or will not appear in the output if there is some ambiguity about it.
// - If an outlying chunk does not have enough elements to reach the 2^abs(n) requirement, it will contain as many elements as remain.
// - Chunks in mirrored places (compared to the median chunk) should be of the same length.
// - You must maintain original element order within each chunk.
// - You must maintain original element order from chunk to chunk.

package kata

func MirroredExponentialChunks(arr []int) [][]int {
	// TODO
}