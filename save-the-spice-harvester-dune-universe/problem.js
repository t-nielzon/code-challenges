// Background Story:
// In the Dune universe, a spice harvester is a large, mobile factory designed to harvest
// the spice Melange. Harvesters are regularly eaten by sandworms, so spotters deploy
// ornithopters to watch for wormsign and contact the nearest Carryall for pickup.
//
// Goal:
// Determine whether a carryall should be sent for rescue, or if it must be forfeited
// because there is not enough time.
//
// Input: object with:
// - harvester: location [x, y]
// - worm: [[location], speed]
// - carryall: [[location], speed]
//
// Conditions:
// - All coordinates are [x, y] (may be positive or negative)
// - Worm and Carryall move toward harvester in straight lines at constant speed
// - Carryall takes 1 minute to lift harvester to safe altitude
// - Distance in kilometers, speed in km/minute
// - Do not mutate input
//
// Output:
// If harvester can be saved: 'The spice must flow! Rescue the harvester!'
// Otherwise: 'Damn the spice! I'll rescue the miners!'

function harvesterRescue(data) {
  // Your solution here
}