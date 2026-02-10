/*
 * Readability is King
 * 5 kyu
 *
 * We will use the Flesch–Kincaid Grade Level to evaluate the readability
 * of a piece of text. This grade level is an approximation for what
 * schoolchildren are able to understand a piece of text.
 *
 * The way to calculate the grade level is as follows:
 *
 *   (0.39 * average number of words per sentence) +
 *   (11.8 * average number of syllables per word) - 15.59
 *
 * Write a function that will calculate the Flesch–Kincaid grade level
 * for any given string. Return the grade level rounded to two decimal points.
 *
 * Ignore hyphens, dashes, apostrophes, parentheses, ellipses and abbreviations.
 *
 * HINT: Count the number of vowels as an approximation for the number of
 * syllables, but count groups of vowels as one (e.g. deal is one syllable).
 * Do not count y as a vowel!
 */

function fleschKincaid(text) {
}

module.exports = fleschKincaid;