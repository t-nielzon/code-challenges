/*
String subpattern recognition II

Similarly to the previous kata, you will need to return a boolean value if the base string 
can be expressed as the repetition of one subpattern.

Changes from the previous kata:
- If a subpattern has been used, it will be present at least twice, meaning the subpattern 
  has to be shorter than the original string.
- The strings you will be given might or might not be created repeating a given subpattern, 
  then shuffling the result.

Examples:
"a"    --> false //no repeated shorter sub-pattern, just one character
"aaaa" --> true  //just one character repeated
"abcd" --> false //no repetitions
"babababababababa" --> true //repeated "ba"
"bbabbaaabbaaaabb" --> true //same as above, just shuffled

Notes:
- Strings will never be empty and can be composed of any character
- Upper and lowercase letters are considered different entities
- Strings can be pretty long (keep an eye on performances!)
*/

function hasSubpattern(str) {
  // Your solution here
}