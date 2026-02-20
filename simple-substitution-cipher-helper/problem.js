/*
 * Simple Substitution Cipher Helper
 *
 * A simple substitution cipher replaces one character from an alphabet with
 * a character from an alternate alphabet, where each character's position
 * in an alphabet is mapped to the alternate alphabet for encoding or decoding.
 *
 * E.g.
 *   const abc1 = "abcdefghijklmnopqrstuvwxyz";
 *   const abc2 = "etaoinshrdlucmfwypvbgkjqxz";
 *
 *   const sub = new SubstitutionCipher(abc1, abc2);
 *   sub.encode("abc") // => "eta"
 *   sub.encode("xyz") // => "qxz"
 *   sub.encode("aeiou") // => "eirfg"
 *
 *   sub.decode("eta") // => "abc"
 *   sub.decode("qxz") // => "xyz"
 *   sub.decode("eirfg") // => "aeiou"
 *
 * If a character provided is not in the opposing alphabet, simply leave it as be.
 */

function SubstitutionCipher(abc1, abc2) {

}

SubstitutionCipher.prototype.encode = function (str) {

};

SubstitutionCipher.prototype.decode = function (str) {

};