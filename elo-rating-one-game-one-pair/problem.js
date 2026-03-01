/*
 * Elo rating - one game, one pair
 *
 * Calculate a chess player's new Elo rating using their previous ratings
 * (experience), their opponent's rating (opponent), the outcome of the
 * new game (score), and the Factor k function (k).
 *
 * Arguments:
 *   experience - array of the player's rating history (last item = current rating).
 *                If empty, use 1000.
 *   opponent   - opponent's rating (number)
 *   score      - 1 (win), 0 (loss), or 0.5 (tie)
 *   k          - optional function taking experience, returns k-factor
 *
 * Default k function (pre-2011 FIDE):
 *   < 30 games        -> 25
 *   never reached 2400 -> 15
 *   reached 2400+      -> 10
 *
 * Formulas:
 *   expectation = 1 / (1 + 10^((opponent - player) / 400))
 *   new_rating  = player + k(experience) * (score - expectation)
 *
 * Returns new rating rounded to nearest integer.
 */
function elo(experience, opponent, score, k) {
}