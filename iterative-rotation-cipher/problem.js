/*
 * Iterative Rotation Cipher (IRC)
 * 
 * Create an object with two methods: encode and decode
 * 
 * Encoding (n rotations):
 * Step 1: Remove all spaces, remember their positions
 * Step 2: Rotate the string right by n characters
 * Step 3: Put spaces back in their original positions
 * Step 4: Rotate each space-separated substring right by n characters
 * Repeat steps 1-4 n times total
 * Prepend n to the result with a space
 * 
 * Decoding:
 * Reverse the encoding process by reversing each iteration in reverse order
 */

const IterativeRotationCipher = {
  encode(n, str) {
    
  },
  
  decode(str) {
    
  }
};