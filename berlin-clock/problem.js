/*
 * The "Berlin Clock" is the first public clock in the world that tells the time
 * by means of illuminated, coloured fields, for which it entered the Guinness Book
 * of Records upon its installation on 17 June 1975.
 * 
 * The clock is read from the top row to the bottom:
 * - The round Yellow light on top blinks to denote even- (when on) or odd-numbered
 *   (off) seconds.
 * - The first row of four Red fields denote five full hours each.
 * - The second row, also of four Red fields, denote one full hour each (displaying
 *   the hour value in 24-hour format).
 * - The third row consists of eleven Yellow-and-Red fields, which denote five full
 *   minutes each (the red ones also denoting 15, 30 and 45 minutes past).
 * - The bottom row has another four Yellow fields, which mark one full minute each.
 * 
 * Complete the function that takes a particular time in 24h format ("hh:mm:ss")
 * and outputs a string that reproduces the Berlin Clock. The lights should be
 * represented as follows:
 * - R: Red
 * - Y: Yellow
 * - O: Off
 * 
 * The rows of the output string should be joined with newlines (\n).
 * 
 * Examples:
 * "12:56:01" => "O\nRROO\nRROO\nYYRYYRYYRYY\nYOOO"
 * "00:00:00" => "Y\nOOOO\nOOOO\nOOOOOOOOOOO\nOOOO"
 * "22:32:45" => "O\nRRRR\nRROO\nYYRYYROOOOO\nYYOO"
 */

function berlinClock(timeStr) {
  
}