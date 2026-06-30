/*
 * Bitwise logical negation
 *
 * Task
 * Given a number `n`, return the logical negation of `n`, i.e., return `1`
 * if `n = 0`, otherwise return `0`.
 *
 * Restrictions
 *   - You may only use at most 28 characters
 *   - You may only use the following characters: [0-9], [a-z], (, ), +, ~,
 *     &, |, ^, <, =, > ([A-Z] and whitespace are not allowed)
 *   - You may only use the character `+` at most 2 times
 *   - You may not use the following operators: &&, &&=, ||, ||=, <, <=, ==,
 *     ===, >=, >. Only +, ~, &, |, ^, <<, >>, >>> (and their assignment
 *     counterparts) are allowed.
 *
 * Examples
 *   negation(0)  // 1
 *   negation(1)  // 0
 *   negation(-1) // 0
 *   negation(32) // 0
 *
 * Constraints
 *   -2147483648 <= n <= 2147483647
 */

negation=n=>