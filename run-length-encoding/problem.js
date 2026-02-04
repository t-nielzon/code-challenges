/*
 * Run-length encoding
 * 
 * Run-length encoding (RLE) is a very simple form of data compression in which
 * runs of data (that is, sequences in which the same data value occurs in many
 * consecutive data elements) are stored as a single data value and count,
 * rather than as the original run.
 * 
 * Task:
 * Your task is to write such a run-length encoding. For a given string, return
 * a list (or array) of pairs (or arrays) [[i1, s1], [i2, s2], …, [in, sn]],
 * such that one can reconstruct the original string by replicating the character
 * sx ix times and concatenating all those strings. Your run-length encoding
 * should be minimal, ie. for all i the values si and si+1 should differ.
 * 
 * Examples:
 * runLengthEncoding("hello world!")
 *  //=> [[1,'h'], [1,'e'], [2,'l'], [1,'o'], [1,' '], [1,'w'], [1,'o'], [1,'r'], [1,'l'], [1,'d'], [1,'!']]
 * 
 * runLengthEncoding("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabbb")
 *  // => [[34,'a'], [3,'b']]
 */

function runLengthEncoding(str) {
  // Your solution here
}