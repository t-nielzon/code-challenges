/*
 * Run-length encoding (RLE) is a very simple form of data compression
 * in which runs of data are stored as a single data value and count,
 * rather than as the original run.
 *
 * For a given string, return an array of pairs [count, character],
 * such that one can reconstruct the original string by replicating
 * the character count times and concatenating all those strings.
 * The encoding should be minimal, i.e. consecutive pairs must have
 * different characters.
 *
 * Examples:
 *   runLengthEncoding("hello world!")
 *     //=> [[1,'h'],[1,'e'],[2,'l'],[1,'o'],[1,' '],[1,'w'],[1,'o'],[1,'r'],[1,'l'],[1,'d'],[1,'!']]
 *
 *   runLengthEncoding("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabbb")
 *     //=> [[34,'a'],[3,'b']]
 */
function runLengthEncoding(str) {
}