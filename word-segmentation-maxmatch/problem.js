/*
 * Word Segmentation: MaxMatch
 *
 * Some languages like Chinese, Japanese, and Thai do not have spaces between
 * words. However, most natural languages processing tasks like part-of-speech
 * tagging require texts that have segmented words. A simple and reasonably
 * effective algorithm to segment a sentence into its component words is called
 * "MaxMatch".
 *
 * MaxMatch starts at the first character of a sentence and tries to find the
 * longest valid word starting from that character. If no word is found, the
 * first character is deemed the longest "word", regardless of its validity.
 * In order to find the rest of the words, MaxMatch is then recursively invoked
 * on all of the remaining characters until no characters remain. A list of all
 * of the words that were found is returned.
 *
 * Write maxMatch, which takes an alphanumeric, spaceless, lowercased String as
 * input and returns an Array of all the words found, in the order they were
 * found. All valid words are in the Set VALID_WORDS.
 */

function maxMatch(sentence) {
}