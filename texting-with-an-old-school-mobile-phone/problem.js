/*
Texting with an old-school mobile phone (6 kyu)

Key layout:
-------------------------
|   1   |   2   |   3   |  <-- hold a key to type a number
|  .,?! |  abc  |  def  |  <-- press a key to type a letter
-------------------------
|   4   |   5   |   6   |
|  ghi  |  jkl  |  mno  |
-------------------------
|   7   |   8   |   9   |
|  pqrs |  tuv  |  wxyz |
-------------------------
|   *   |   0   |   #   |  <-- hold for *, 0 or #
|  '-+= | space |  case |  <-- press # to switch upper/lower case
-------------------------

Rules:
- Pressing a key n times: cycles to nth char on bottom row.
- Holding (digit + '-'): types the top-row symbol.
- '#' press toggles case (start lowercase). Only consider switch
  when next char is alphabetic.
- If two consecutive characters share the same button, insert a
  space to wait. Exceptions: no wait after holding, no wait when
  a case switch happens between the two characters.

Return the sequence of key inputs (top row) producing `message`
with the lowest number of clicks.
*/

function sendMessage(message) {
  // your code here
}