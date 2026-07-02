// Braking distance d1 is the distance a vehicle will go from the point when it
// brakes to when it comes to a complete stop. It depends on the original speed v
// and on the coefficient of friction mu between the tires and the road surface.
//
// The braking distance is one of two principal components of the total stopping
// distance. The other component is the reaction distance, which is the product of
// the speed and the perception-reaction time of the driver.
//
// The kinetic energy E is 0.5*m*v**2, the work W given by braking is mu*m*g*d1.
// Equalling E and W gives the braking distance:
//   d1 = v*v / (2*mu*g)
// where g is the gravity of Earth and m the vehicle's mass.
//
// v is in km per hour, g is 9.81 m/s/s, and the reaction time is constant and
// equal to 1 s. The coefficient mu is dimensionless.
//
// Tasks:
//   - dist(v, mu)  -> total stopping distance in meters (reaction time t = 1)
//   - speed(d, mu) -> v in km/h such that dist(v, mu) = d (reaction time t = 1)
//
// Examples:
//   dist(100, 0.7)                  -> 83.9598760937531
//   speed(83.9598760937531, 0.7)    -> 100.0
//
// Notes:
//   - Convert velocity from km/h to m/s or from m/s to km/h when necessary.
//   - Don't forget the reaction time t = 1.
//   - Don't truncate or round your results.

package kata

func dist(v, mu float64) float64 {
	// your code here
}

func speed(d, mu float64) float64 {
	// your code here
}