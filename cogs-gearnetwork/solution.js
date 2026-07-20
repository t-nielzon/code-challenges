function cogsebi(gears, connections, driver_id, driver_rpm) {
  const n = gears.length;
  const rpms = new Array(n).fill(0);
  
  // Build adjacency list from connections
  const adj = Array.from({ length: n }, () => []);
  for (const [a, b] of connections) {
    adj[a].push(b);
    adj[b].push(a);
  }
  
  // BFS from driver gear to propagate RPMs through the network
  const queue = [driver_id];
  rpms[driver_id] = driver_rpm;
  const visited = new Set([driver_id]);
  
  while (queue.length > 0) {
    const current = queue.shift();
    const currentRpm = rpms[current];
    const currentTeeth = gears[current];
    
    for (const neighbor of adj[current]) {
      if (!visited.has(neighbor)) {
        visited.add(neighbor);
        const neighborTeeth = gears[neighbor];
        // gear direction reverses when driven by another, ratio determined by teeth
        rpms[neighbor] = -currentRpm * (currentTeeth / neighborTeeth);
        queue.push(neighbor);
      }
    }
  }
  
  return rpms;
}