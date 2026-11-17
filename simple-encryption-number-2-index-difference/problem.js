/*
 * Simple Encryption #2 - Index-Difference
 * 
 * Encryption region (77 chars):
 * - Uppercase letters A-Z (0-25)
 * - Lowercase letters a-z (26-51)
 * - Digits 0-9 (52-61)
 * - Special chars: .,:;-?! '()$%&" (62-76)
 * 
 * Encryption steps:
 * 1. For every second char (positions 1,3,5...), switch case
 * 2. For each char (starting at position 1), replace with char at index
 *    (previous_char_index - current_char_index) mod 77. First char unchanged.
 * 3. Replace first char with its mirror: index 76 - original_index
 * 
 * Decryption: Reverse the steps above
 */

function encrypt(text) {
  // placeholder
}

function decrypt(encryptedText) {
  // placeholder
}