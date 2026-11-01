/*
 * Things like blog posts and CMS pages have titles. But they're accessed through urls.
 * Titles can have old character in them. They can even be exactly the same as another one.
 * So the back-end has to simplify those titles to be valid for a url at all, while still
 * making sure they stay unique.
 *
 * Write the function `simunq` that converts the given original string and returns it with
 * all the weird, disallowed characters replaced by a "wildcard."
 *
 * Allowed characters: all english lowercase and uppercase characters, all digits,
 * the space character, and the wildcard (dash `-`).
 *
 * Example: simunq("blah?") == "blah-";
 *
 * But at the same time, two urls must never overlap. So if it's the same as an
 * already-used simunq string, tag a `-n` to the end of it, where `n` is the number
 * of identical strings that came before it, plus 1 for the current string.
 *
 * Examples:
 * simunq("blah") == "blah";
 * simunq("blah") == "blah-2";
 * simunq("blah-2") == "blah-2-2";
 *
 * Add a `reset` method to the `simunq` function. After this has been called,
 * all the previously created simunq strings should be ignored.
 *
 * Examples:
 * simunq("string") == "string";
 * simunq("string") == "string-2";
 * simunq.reset();
 * simunq("string") == "string";
 */

function simunq(str) {
  // Your code here
}

simunq.reset = function() {
  // Your code here
};