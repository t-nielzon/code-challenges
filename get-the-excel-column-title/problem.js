/*
 * Get the Excel column title!
 *
 * Implement the getColumnTitle(num) function that takes an integer number
 * (index of the Excel column) and returns the string representing the title
 * of this column.
 *
 * The first column has the title "A", second - "B", 26th - "Z", 27th - "AA".
 * "BA"(53) comes after "AZ"(52), "AAA" comes after "ZZ".
 *
 * Input: column decimal index number (natural number)
 * Output: upper-case string representing the column title (A..Z)
 *
 * Errors:
 * - For num < 1: throw IndexError
 * - For non-integer argument: throw TypeError
 *
 * Examples:
 *   getColumnTitle(52)     => "AZ"
 *   getColumnTitle(1337)   => "AYK"
 *   getColumnTitle(432778) => "XPEH"
 */

function getColumnTitle(num) {
}

module.exports = { getColumnTitle };