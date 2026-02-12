/*
 * Four Letter Words ~ Anagrams
 *
 * Alice and Bob (Team One) vs Carol and Dan (Team Two) play an anagram game.
 * 
 * For each subject word, two players compete in a round:
 * - Both get 1 point if they know the subject word itself
 * - Player One (first) gets 2 points per anagram they recall
 * - Player Two (second) gets 3 points per anagram Player One didn't recall
 * 
 * The player with the higher round score wins 1 final point for their team.
 * Ties award no final points.
 *
 * The Order of Rounds:
 *   1. Alice vs Carol    5. Alice vs Dan
 *   2. Carol vs Bob      6. Carol vs Alice
 *   3. Bob vs Dan        7. Bob vs Carol
 *   4. Dan vs Alice      8. Dan vs Bob
 *
 * Players: Alice=0, Bob=1, Carol=2, Dan=3
 * Team One: Alice(0), Bob(1) | Team Two: Carol(2), Dan(3)
 *
 * Input:
 *   subjects - array of four-letter words
 *   memories - array of 4 arrays (one per player) of four-letter words
 *   players  - array of 2 integers, the starting pair
 *
 * Output:
 *   0 if Team One wins, 1 if Team Two wins, -1 if tie
 */
function anagrams(subjects, memories, players) {

}

module.exports = { anagrams };