/*
 * Statistics for an Athletic Association
 *
 * You are the "computer expert" of a local Athletic Association (C.A.A.).
 * Many teams of runners come to compete. Each time you get a string of
 * all race results of every team who has run.
 *
 * For example here is a string showing the individual results of a team of 5 runners:
 *   "01|15|59, 1|47|6, 01|17|20, 1|32|34, 2|3|17"
 *
 * Each part of the string is of the form: h|m|s
 * where h, m, s (h for hour, m for minutes, s for seconds) are positive or null
 * integer (represented as strings) with one or two digits. Substrings in the
 * input string are separated by ", " or ",".
 *
 * To compare the results of the teams you are asked for giving three statistics:
 * range, average and median.
 *
 * - Range: difference between the lowest and highest values.
 * - Mean or Average: sum of all values divided by the total count.
 * - Median: middle value of the sorted list (or average of the two middle values
 *   if the count is even).
 *
 * Return a string of the form:
 *   "Range: hh|mm|ss Average: hh|mm|ss Median: hh|mm|ss"
 *
 * Remarks:
 *   1. If a result in seconds is ab.xy... it will be given truncated as ab.
 *   2. If the given string is "" you will return "".
 */

function stat(strg) {

}