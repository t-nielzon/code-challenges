/*
 * Clone-a-Regex
 * Difficulty: 6 kyu
 *
 * Make a utility function to clone a regex object that will give exactly the
 * same results as the original. The functionality should be able to be used
 * in two methods: `RegExp.clone(/regex/)` and `(/regex/).clone()`. Just
 * because. ;P
 *
 * NOTE: any time a regex is "run" (using match, replace, exec, etc.), its
 * state may change such that the next time it is run, a different result is
 * produced.
 */

RegExp.clone = function (regex) {
  // your code here
};

RegExp.prototype.clone = function () {
  // your code here
};