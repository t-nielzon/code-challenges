/*
 * Four Letter Words ~ Anagrams
 *
 * Alice and Bob (Team One) vs Carol and Dan (Team Two) play an anagram game.
 *
 * Given a list of four-letter subject words, player memories, and starting players,
 * determine the winner.
 *
 * Each round:
 * 1. Both players get 1 point if they know the subject word itself.
 * 2. Player One (first) recalls anagrams for 2 points each.
 * 3. Player Two recalls any remaining anagrams for 3 points each.
 *
 * The round winner earns 1 final point for their team.
 *
 * The Order of Rounds:
 * 1. Alice vs Carol    5. Alice vs Dan
 * 2. Carol vs Bob      6. Carol vs Alice
 * 3. Bob vs Dan        7. Bob vs Carol
 * 4. Dan vs Alice      8. Dan vs Bob
 *
 * Input:
 *   subjects - array of four-letter words
 *   memories - array of 4 arrays (Alice=0, Bob=1, Carol=2, Dan=3)
 *   players  - array of 2 integers, the starting pair
 *
 * Output:
 *   0 if Team One wins, 1 if Team Two wins, -1 if tie
 */

function anagrams(subjects, memories, players) {
}