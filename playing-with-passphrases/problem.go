/*
Playing with passphrases

Everyone knows passphrases. One can choose passphrases from poems, songs, movies names and so on but frequently
they can be guessed due to common cultural references.
You can get your passphrases stronger by different means. One is the following:

choose a text in capital letters including or not digits and non alphabetic characters,

1. shift each letter by a given number but the transformed letter must be a letter (circular shift),
2. replace each digit by its complement to 9,
3. keep such as non alphabetic and non digit characters,
4. downcase each letter in odd position, upcase each letter in even position (the first character is in position 0),
5. reverse the whole result.
*/
package kata

func PlayPass(s string, n int) string {
}