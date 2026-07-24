/*
Guess Who? - A two-player guessing game

Your task is to make a simple class called GuessWho. The computer will try and guess your character,
your job is to return back to the computer a list of possible characters so that it can guess successfully.

The computer will:
1. Give you the character in the initialization of the class
2. Post the guess to the method 'guess'
3. The guess will either be a name of a character or one of these characteristics:
   ["Male","Female","Glasses","Brown eyes","Bald","White hair","Small mouth","Mustache","Brown hair",
    "Big mouth","Small nose","Blue eyes","Hat","Long hair","Black hair","Earrings","Blonde hair",
    "Ginger hair","Beard","Big nose"]

Rules:
- If the characteristic matches your character, return all characters with that characteristic
- If the characteristic doesn't match your character, return all characters without that characteristic
- Update your character list
- Keep a tally of the amount of turns
- Return ["Correct! in n turns"] when the computer guesses correctly
- Return array of possible characters otherwise
*/

class GuessWho {
  constructor(character) {
    // Initialize with the character name
  }

  guess(g) {
    // Process a guess and return results
  }
}