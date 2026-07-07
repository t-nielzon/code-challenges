/*
 * A chain of car rental locals made a statistical research for each local in a city.
 * They measure the average of clients for every day.
 *
 * They used the Simpson model for probabilities:
 * P(x) = (λ^x * e^-λ) / x!
 *
 * where:
 * - λ: average for the amount of events having the same conditions during the measurement
 * - x: number of the events different from the average at the same place and time interval
 * - e: Euler's number (2.7182818284...)
 *
 * Make the function prob_simpson() that will receive 2 or 3 arguments.
 *
 * 1) Two arguments if we want to calculate the probability for a specific number,
 *    receives the average lamb and the variable x.
 *
 * 2) Three arguments if we want the cumulative probability for a value, besides receiving
 *    lamb and x, will receive one of the following operators in string format: <, <=, >, >=
 *
 * Examples:
 * prob_simpson(8, 12) == 0.04812680428195667
 * prob_simpson(8, 12, '>') == 0.888075998981
 * prob_simpson(8, 12, '>=') == 0.936202803263
 * prob_simpson(8, 12, '<') == 0.063797196737
 * prob_simpson(8, 12, '<=') == 0.111924001019
 */

function prob_simpson(lamb, x, operator) {
  
}