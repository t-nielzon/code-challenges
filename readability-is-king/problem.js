/*
 * Readability is King
 *
 * We will use the Flesch–Kincaid Grade Level to evaluate the readability
 * of a piece of text. This grade level is an approximation for what
 * schoolchildren are able to understand a piece of text.
 *
 * Formula:
 *   (0.39 * average number of words per sentence) +
 *   (11.8 * average number of syllables per word) - 15.59
 *
 * Return the grade level rounded to two decimal points.
 *
 * Ignore hyphens, dashes, apostrophes, parentheses, ellipses and abbreviations.
 *
 * HINT: Count vowels as an approximation for syllables, but count groups
 * of vowels as one. Do not count y as a vowel.
 */
function fleschKincaid(text) {
}

module.exports = { fleschKincaid };