/*
 * Binary Multiplication - Ancient Egyptian Method
 * 
 * The goal of this kata is to multiply two integers using the ancient Egyptian method,
 * which only requires divisions and multiplications by two, and additions.
 * 
 * Your function takes two integers as input. It shall return a list of the steps in
 * the multiplication.
 * 
 * Let m be the largest and n the smallest. While m is superior to 0, at each step:
 * - if m is not divisible by 2, add n to the list
 * - divide m by 2 (integer division)
 * - multiply n by 2
 * 
 * At the end, return the list in descending order.
 * The result of the multiplication is the sum of the list elements, but you have to
 * return only the list.
 * 
 * Input: n and m integers, from 0 to 10,000. It is not guaranteed that m > n.
 * 
 * Example:
 * egyptianMultiplication(100, 15) returns [960, 480, 60]
 * because 100 * 15 = 1500 = 960 + 480 + 60
 */

function egyptianMultiplication(a, b) {
  
}