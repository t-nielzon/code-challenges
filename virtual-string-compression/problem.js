/*
 * Run-Length Encoding (RLE) for Virtual String
 * 
 * You are given a single number, length, representing the length of the string.
 * You cannot access the string directly; use getChar(position) to get the 
 * character at a 0-based index.
 * 
 * Return the RLE of the string: each consecutive sequence of the same character 
 * is represented as <count><char>.
 * 
 * Blocks such as 0w, 00a, or any block starting with a zero are not valid RLE 
 * and must not appear in the output.
 * 
 * All characters in the virtual string will be non-digit characters, so no output 
 * block will ever contain a digit as the character itself.
 * 
 * It is guaranteed that no character will appear in more than one block — each 
 * character block is unique.
 * 
 * Examples:
 * - If the string were "aaabbc", return "3a2b1c"
 * - If the string were "xxyz", return "2x1y1z"
 * 
 * The string length will never exceed 10^15 and never be smaller than 1.
 */

function rle(length, getChar) {
  // Your solution here
}