/*
The "Berlin Clock" is the first public clock in the world that tells the time by means of illuminated, coloured fields.

The clock is read from the top row to the bottom:
- the round Yellow light on top blinks to denote even- (when on) or odd-numbered (off) seconds,
- the first row of four Red fields denote five full hours each,
- the second row, also of four Red fields, which denote one full hour each (displaying the hour value in 24-hour format)
- the third row consists of eleven Yellow-and-Red fields, which denote five full minutes each (the red ones also denoting 15, 30 and 45 minutes past),
- the bottom row has another four Yellow fields, which mark one full minute each.

Complete the function that takes a particular time in 24h format ("hh:mm:ss") and outputs a string that reproduces the Berlin Clock.
Lights should be represented as:
* R : Red
* Y : Yellow
* O : Off

The rows of the output string should be joined with newlines (\n).

Examples:
"12:56:01"  ==>  "O\nRROO\nRROO\nYYRYYRYYRYY\nYOOO"
"00:00:00"  ==>  "Y\nOOOO\nOOOO\nOOOOOOOOOOO\nOOOO"
"22:32:45"  ==>  "O\nRRRR\nRROO\nYYRYYROOOOO\nYYOO"
*/

function berlinClock(timeString) {
  
}