/**
 * Desert Race (6 kyu)
 * 
 * The desert festival is hosting a camel race.
 * 
 * Each camel starts at position 0.
 * 
 * You are given race data containing:
 * - a collection of registered camels (each with id and name)
 * - a collection of race events, in the order they happened
 * 
 * Your task is to process all events and return the camel, or camels, that
 * finished the race in the farthest position.
 * 
 * Event types:
 * - move: changes position by value (positive = forward, negative = backward)
 * - sandstorm: moves camel back 2 positions
 * - oasis: makes the next move event for that camel count double
 *   (only affects the next move, multiple oasis before move count only once)
 * 
 * Unknown camel IDs in events should be ignored.
 * 
 * Return an array containing the name(s) of the camel(s) that finished farthest.
 * If multiple camels tie, return their names sorted alphabetically.
 */
function desertRace(camels, events) {
  
}