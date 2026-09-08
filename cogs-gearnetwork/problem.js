/*
A series of cogs are connected in a sprawling network. Unlike simple gear trains arranged 
in a rigid straight line, a single cog in this system can drive multiple other cogs 
simultaneously, creating complex branches and webs of mechanical power!

You are given a list of cogs, their tooth counts, and a list of physical connections. 
One specific cog is the driver cog, which spins at a given Revolutions Per Minute (RPM). 
Your job is to calculate the final RPM of all cogs in the network.

Rules of Gear Engagement:
1. Speed Ratio: If Cog A drives Cog B, the speed ratio is: RPM_B = RPM_A * (Teeth_A / Teeth_B)
2. Direction (Rotation): Every time a gear drives another, the direction reverses.
   - Clockwise rotation is represented by a positive RPM.
   - Counter-clockwise rotation is represented by a negative RPM.
3. The Driver: The driver cog's starting RPM is always given as a positive number (clockwise).
4. Valid Networks: All provided gear networks are physically valid. No locks or jams.

Input:
- gears: An array of integers where the index is the Gear ID, and the value is the 
  number of teeth.
- connections: An array of pairs [A, B] representing a physical mesh between Gear A and B.
- driver_id: The integer ID of the gear providing the power.
- driver_rpm: The starting RPM of the driver gear (can be an integer or float).

Output:
- Return an array of RPMs for all gears (ordered by their ID).
- Do not round the final results. Return the exact floating-point values.
- If a gear is completely disconnected from the driver system, its RPM is 0.0.
*/

function cogsebi(gears, connections, driver_id, driver_rpm) {
  
}