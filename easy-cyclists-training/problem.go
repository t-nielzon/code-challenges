/*
Easy Cyclist's Training

John has bought a bike but before going mountain biking he wants us to do a few simulations.

He gathered information:
- His trip will consist of an ascent of dTot kilometers with an average slope of slope percent
- We suppose that: there is no wind, John's mass is constant MASS = 80 (kg), his power
  (generated at time t by pedaling and measured in watts) at the start of the trip is WATTS0 = 225 (watts)
- We don't take account of the rolling resistance
- When he starts climbing at t = 0 his initial speed (pushed start) is v0 (km/h)
- His initial acceleration gamma is 0.
- Our time step is DELTA_T = 1.0 / 60.0 (in minutes)

Write function temps(v0, slope, dTot) which returns as a rounded integer the time t in minutes
needed to travel dTot. If John gives up return -1.
*/
package kata

func Temps(v0, slope, dTot float64) int {
}