/*
 * HTML Complementary Color
 *
 * Implement a function that takes a hex-color string and returns the string
 * representing the complementary color.
 *
 * Input:
 * A hexadecimal string (case-ignored with chars 0..9 or A..F) without hash "#".
 * If shorter than 6 characters, pad with zeroes from the left to make it 6 chars.
 *   "a23" <=> "000a23"
 *   "" <=> "0" <=> "000000"
 *
 * Output:
 * An uppercased string containing "#" followed by the complementary color.
 * Complementary color gives white when summed with the input:
 *   #000A23 + #FFF5DC = #FFFFFF
 *
 * Errors:
 * Throw an Error if: string length is 7+, has non-hex chars, or non-string type.
 *
 * Examples:
 *   "01fD08" --> "#FE02F7"
 *      "a23" --> "#FFF5DC"
 *         "" --> "#FFFFFF"
 */
function complementaryColor(hex) {
}

module.exports = { complementaryColor };