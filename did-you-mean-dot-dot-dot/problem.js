/*
 * Did you mean ...? (5 kyu)
 *
 * You'll get an entered term (lowercase string) and an array of known words
 * (also lowercase strings). Your task is to find out which word from the
 * dictionary is most similar to the entered one. The similarity is described
 * by the minimum number of letters you have to add, remove or replace in
 * order to get from the entered word to one of the dictionary (Levenshtein distance).
 */

function Dictionary(words) {
  this.words = words;
}

Dictionary.prototype.findMostSimilar = function(term) {

};