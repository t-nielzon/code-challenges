/*
 * Berlin Clock
 * Difficulty: 6 kyu
 * 
 * The "Berlin Clock" is the first public clock in the world that tells the time by means 
 * of illuminated, coloured fields. It reads from top to bottom:
 * 
 * - The round Yellow light on top blinks to denote even- (when on) or odd-numbered (off) seconds
 * - The first row of four Red fields denote five full hours each
 * - The second row of four Red fields denote one full hour each (24-hour format)
 * - The third row of eleven Yellow-and-Red fields denote five full minutes each 
 *   (the red ones also denoting 15, 30 and 45 minutes past)
 * - The bottom row has four Yellow fields, which mark one full minute each
 * 
 * Complete the function that takes a time in 24h format ("hh:mm:ss") and outputs 
 * a string that reproduces the Berlin Clock.
 * 
 * Lights are represented as:
 * - R: Red
 * - Y: Yellow
 * - O: Off
 * 
 * Rows should be joined with newlines (\n).
 * 
 * Examples:
 * "12:56:01" => "O\nRROO\nRROO\nYYRYYRYYRYY\nYOOO"
 * "00:00:00" => "Y\nOOOO\nOOOO\nOOOOOOOOOOO\nOOOO"
 * "22:32:45" => "O\nRRRR\nRROO\nYYRYYROOOOO\nYYOO"
 */

function berlinClock(time) {
  
}