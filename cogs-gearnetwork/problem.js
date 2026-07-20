/**
 * Cogs Gear Network
 * 
 * A series of cogs are connected in a sprawling network. Unlike simple gear trains
 * arranged in a rigid straight line, a single cog in this system can drive multiple
 * other cogs simultaneously, creating complex branches and webs of mechanical power!
 * 
 * You are given a list of cogs, their tooth counts, and a list of physical connections.
 * One specific cog is the driver cog, which spins at a given RPM. Calculate the final
 * RPM of all cogs in the network.
 * 
 * Rules:
 * 1. Speed Ratio: RPM_B = RPM_A * (Teeth_A / Teeth_B)
 * 2. Direction: Every time a gear drives another, direction reverses.
 *    Positive RPM = Clockwise, Negative RPM = Counter-clockwise
 * 3. Valid Networks: No locks or jams, no conflicting ratios
 * 
 * @param {number[]} gears - Array where index is Gear ID, value is tooth count
 * @param {number[][]} connections - Array of pairs [A, B] representing meshed gears
 * @param {number} driver_id - ID of the driver gear
 * @param {number} driver_rpm - Starting RPM of the driver gear (positive)
 * @returns {number[]} Array of RPMs for all gears, ordered by ID (0.0 if disconnected)
 */
function cogsebi(gears, connections, driver_id, driver_rpm) {
  
}