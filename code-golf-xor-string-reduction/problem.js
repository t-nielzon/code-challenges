/*
 * [Code Golf] XOR string reduction
 *
 * Given a string consisting entirely of binary digits (0, 1) separated
 * using spaces. Find the XOR of all digits and return the answer.
 *
 * Examples:
 *   "1 0 0 1 0" --> 0
 *   "1 0 1 1 1 0 0 1 0 0 0 0" --> 1
 *
 * How:
 *   1 0 0 1 0
 *   (1 XOR 0) (0 XOR 1) 0
 *   1 1 0
 *   (1 XOR 1) 0
 *   0 0
 *   0 XOR 0
 *   0 ---> Answer
 *
 * This is code-golf so shortest code wins. Limit: 40 chars (exclusive).
 */
xor=s=>{}