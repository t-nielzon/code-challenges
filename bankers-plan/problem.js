/*
 * Banker's Plan
 *
 * John has some amount of money of which he wants to deposit a part f0 to the bank at the beginning
 * of year 1. He wants to withdraw each year for his living an amount c0.
 *
 * - deposit f0 at beginning of year 1
 * - his bank account has an interest rate of p percent per year, constant over the years
 * - John can withdraw each year c0, taking it whenever he wants in the year; he must take account
 *   of an inflation of i percent per year in order to keep his quality of living.
 * - all amounts f0..fn-1, c0..cn-1 are truncated by the bank to their integral part
 * - Given f0, p, c0, i the banker guarantees that John will be able to go on that way until the nth year.
 *
 * Given parameters f0, p, c0, n, i build a function fortune which returns true if John can make
 * a living until the nth year and false if it is not possible.
 */
function fortune(f0, p, c0, n, i) {
}

module.exports = { fortune };