/*
In this Kata, we will be taking the first step in solving Nash Equilibria of 2 player games
by learning to calculate the Expected Utility of a player in a game of Rock Paper Scissors
given the Utilities of the game and the Pure/Mixed Strategy of each Player.

Let's first define a few Game Theoretic Definitions:

S_i = finite set of actions or choices for player i.
A = S_1 x ... x S_n is the set of all possible combinations of actions of all players.
    Each possible combination of actions is called an action profile.
sigma_i(s) = Probability player i chooses action s in S_i
u = a function mapping each action profile to a vector of utilities for each player.

We refer to player i's payoff as u_i and by convention, -i refers to player i's opponents.

To compute the expected utility of the game for a player, sum over each action profile
the product of each player's probability of playing their action in the action profile,
times the player's utility for the action profile:

    u_i(sigma_i, sigma_-i) = sum over s in S_i, s' in S_-i of
                             sigma_i(s) * sigma_-i(s') * u_i(s, s')

The payoff table for Rock-Paper-Scissors is:

      |   R   |   P   |   S
    R |  0, 0 | -1, 1 |  1,-1
    P |  1,-1 |  0, 0 | -1, 1
    S | -1, 1 |  1,-1 |  0, 0

Which can be expressed as:
    utilities = [[0,-1,1],[1,0,-1],[-1,1,0]]

Example:
    utilities = [[0,-1,1],[1,0,-1],[-1,1,0]]
    p0 = [1,0,0]
    p1 = [0,1,0]
    expectedUtility(p0, p1, utilities) === -1

The returned value should be rounded to 2 decimal places.
*/

function expectedUtility(p0, p1, utilities) {
  // your code here
}