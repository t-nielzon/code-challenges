/*
 * Mod4 Regex
 *
 * Write a Regular Expression that matches any string with at least one number
 * divisible by 4 (with no remainder).
 *
 * A number will start with [ and end with ]. They may (or may not) include a
 * plus or minus symbol at the start; this should be taken into account. Leading
 * zeros may be present, and should be ignored. There may be other text in the
 * string, outside of the number; this should also be ignored. All numbers will
 * be integers; any floats should be ignored.
 *
 * If there are no valid numbers defined as above, there should be no match.
 *
 * Examples:
 *   "[+05620]" // valid (5620 divisible by 4)
 *   "[+05621]" // invalid
 *   "[-55622]" // invalid
 *   "[005624]" // valid
 *   "[0]"      // valid (0 is divisible by 4)
 *
 * NOTE: Only Mod4.test(str) will be used.
 */

var Mod4;