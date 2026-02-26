/*
Decode the Morse code

In this kata you have to write a simple Morse code decoder. While the Morse
code is now mostly superseded by voice and digital data communication channels,
it still has its use in some applications around the world.

The Morse code encodes every character as a sequence of "dots" and "dashes".
For example, the letter A is coded as ·−, letter Q is coded as −−·−, and
digit 1 is coded as ·−−−−. The Morse code is case-insensitive, traditionally
capital letters are used. When the message is written in Morse code, a single
space is used to separate the character codes and 3 spaces are used to separate
words.

NOTE: Extra spaces before or after the code have no meaning and should be ignored.

Your task is to implement a function that would take the morse code as input
and return a decoded human-readable string.

Example:
  DecodeMorse(".... . -.--   .--- ..- -.. .")
  // should return "HEY JUDE"

NOTE: For coding purposes you have to use ASCII characters . and -, not Unicode characters.

The Morse code table is preloaded for you as MORSE_CODE['.--'].
*/
package kata

func DecodeMorse(morseCode string) string {
}