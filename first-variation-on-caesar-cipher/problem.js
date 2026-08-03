/*
 * The action of a Caesar cipher is to replace each plaintext letter with a different one
 * a fixed number of places up or down the alphabet.
 * 
 * This program performs a variation where the shift increases by 1 for each character.
 * If the shift is initially 1, the first character is shifted by 1, the second by 2, etc.
 * 
 * movingShift(s, shift):
 *   - Encodes a string with increasing shift
 *   - Returns an array of 5 strings (the encoded message split into parts)
 *   - Parts have non-increasing lengths; parts 1-4 are between floor(length/5) and ceil(length/5)
 *   - params: s (string to encode), shift (initial shift value)
 * 
 * demovingShift(arr, shift):
 *   - Decodes an array of strings back to the original
 *   - params: arr (array of 5 strings from movingShift), shift (initial shift value)
 *   - returns: the decoded string
 */

function movingShift(s, shift) {
  
}

function demovingShift(arr, shift) {
  
}