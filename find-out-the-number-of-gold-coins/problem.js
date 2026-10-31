/*
 * Find Out the Number of Gold Coins!
 * Difficulty: 5 kyu
 * 
 * After a long RPG campaign, you come across a guardian of a treasure chest full 
 * of gold coins. The guardian will only deliver the treasure if you guess the 
 * number of gold coins present.
 * 
 * The guardian gives N tips in the form of (a_i, m_i) pairs:
 * - If the coins were divided between m_i people, a_i coins would be left.
 * 
 * Constraints:
 * - All m_i values are pairwise relatively prime (gcd(m_i, m_j) = 1 for i ≠ j)
 * - 0 ≤ a_i < m_i
 * - Maximum 15 tips
 * - m_i values can be less than 10^20
 * 
 * Task: Return the minimum number of gold coins in the chest.
 * 
 * Example: numberOfCoins([[1, 2], [2, 5]]) returns 7
 * - 7 % 2 = 1 ✓
 * - 7 % 5 = 2 ✓
 */

function numberOfCoins(tips) {
  // Your solution here
}