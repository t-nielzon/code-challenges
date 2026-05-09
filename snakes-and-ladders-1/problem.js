/*
## Introduction

Snakes and Ladders is an ancient Indian board game regarded today as a worldwide classic.
It is played by two or more players on a game board with numbered, gridded squares.
A number of "ladders" and "snakes" are pictured on the board, each connecting two specific squares.

## Task

Create a class called `SnakesLadders`. The test cases will call the method `play(die1, die2)`
independently of the state of the game or the player turn. The arguments `die1` and `die2`
are the dice thrown in a turn and are both integers between 1 and 6. The player will make
a number of steps equal to the sum of `die1` and `die2`, moving one square per step.

## Rules

1. There are two players, both start off the board on square 0.
2. Player 1 starts and alternates with player 2.
3. Follow the numbers up the board in order from 1 to 100.
4. If both dice are the same, that player has another turn.
5. Climb up ladders (even on a double).
6. Slide down snakes (even on a double).
7. Land exactly on the last square to win. Bounce off 100 if you roll too high.
8. If a player rolls a double and lands on 100 after all steps, they win and do not take another turn.

## Returns

- "Player n Wins!" when player n lands on 100.
- "Game over!" if a move is attempted after any player has won.
- Otherwise, "Player n is on square x".
*/

class SnakesLadders {
  constructor() {
    
  }
  
  play(die1, die2) {
    
  }
}