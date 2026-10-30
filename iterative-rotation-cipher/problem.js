/*
Iterative Rotation Cipher (IRC)

Implement an object with two methods: encode and decode

Encoding:
1. Remove all spaces (remember positions)
2. Rotate the entire string right by n
3. Put spaces back at original positions
4. Rotate each space-separated substring right by n
5. Repeat steps 1-4 n times total
6. Prepend n to the result

Decoding:
Reverse the encoding process

Test example:
IterativeRotationCipher.encode(10, 'If you wish to make an apple pie from scratch, you must first invent the universe.');
// Returns '10 hu fmo a,ys vi utie mr snehn rni tvte .ysushou teI fwea pmapi apfrok rei tnocsclet'
*/

const IterativeRotationCipher = {
  encode: function(n, str) {
  },
  
  decode: function(str) {
  }
};