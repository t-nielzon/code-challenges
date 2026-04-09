/*
Braking distance d1 is the distance a vehicle will go from the point when it brakes
to when it comes to a complete stop. It depends on the original speed v and on the
coefficient of friction mu between the tires and the road surface.

The braking distance is one of two principal components of the total stopping distance.
The other component is the reaction distance, which is the product of the speed and
the perception-reaction time of the driver.

The kinetic energy E is 0.5*m*v**2, the work W given by braking is mu*m*g*d1.
Equalling E and W gives the braking distance:
  d1 = v*v / (2*mu*g)
where g is the gravity of Earth and m the vehicle's mass.

We have v in km per hour, g as 9.81 m/s/s and the reaction time is constant and equal to 1 s.

Tasks:
  dist(v, mu) -> total stopping distance in meters
  speed(d, mu) -> v in km/h such that dist(v, mu) = d
*/
package kata

func Dist(v, mu float64) float64 {
}

func Speed(d, mu float64) float64 {
}