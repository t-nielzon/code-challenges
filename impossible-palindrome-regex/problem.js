/*
Intro
You are known to your friends as 'The Regex Wizard'. You have solved all variety of hard regex katas, and feel pretty confident that you could solve any more that comes your way. Then one day your friend (a computer scientist) approaches you and says "I know you are good at regex, but I have a challenge that you can't solve." "Impossible" you say, "Tell me what it is, I'll solve it." Your friend smiles slyly, "Alright. Write me a regex which will match palindromic words of any length."

Task
Define a constant `palindromeMatcher`, which will match palindromes of any (nonzero) length when tested with the following `String.prototype` methods:
- .match
- .matchAll
- .replace
- .search
- .split

What's a Palindrome?
For this kata we define a palindrome as a single 'word' of length >= 1, which remains the same when reversed (case-insensitive). A 'word' is a sequence of letters (a-z, A-Z) which is not preceded or succeeded by another letter. Other symbols (eg. !?.) are not a part of a word.

Notes:
- Matches should be global.
- Do not try to modify the aforementioned methods (`String` and `RegExp` have been frozen)
- Matches should not contain any capture groups.
- Some words might be rather large.
- Hint: you may want to read up on Symbols.

Examples:
"Hi! Wow its already noon".match(palindromeMatcher) => ["Wow", "noon"]
"i like pop music lol".replace(palindromeMatcher, "$&!?") => "i!? like pop!? music lol!?"
"hey dad catch".search(palindromeMatcher) => 4
"my sis and i made cake".split(palindromeMatcher) => ["my ", " and ", " made cake"]
*/

const palindromeMatcher = {};