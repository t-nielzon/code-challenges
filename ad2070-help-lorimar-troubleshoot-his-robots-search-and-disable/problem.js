/*
AD2070: Help Lorimar troubleshoot his robots-Search and Disable
Difficulty: 6 kyu

The year is 2070 and intelligent connected machines have replaced humans in most things.
Lorimar is a forward deployed machine engineer tracking an infected, non-responsive BOT robot.

Your task is to search through a string of integers related to BOT system internals
looking for a pattern indicating the BOT should be disabled.

The pattern requires ALL of the following:
- Integer is Prime
- Integer is 4 digits long (e.g., 1031)
- Integer's third digit is either 2 or 3 (e.g., 1021 or 1031)
- The count of the integer in the log is > 3

If there are > 50 total integers in the log matching this pattern, 
the BOT at this location should be disabled.

Return "match disable bot" if count > 50
Return "no match continue" otherwise
*/

function helpLorimar(log) {
  
}