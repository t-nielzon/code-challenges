function cogsebi(gears, connections, driver_id, driver_rpm) {
  const n = gears.length;
  const rpms = new Array(n).fill(0);
  const visited = new Set();
  
  // build adjacency list for the gear network
  const adj = Array.from({ length: n }, () => []);
  for (const [a, b] of connections) {
    adj[a].push(b);
    adj[b].push(a);
  }
  
  // bfs to traverse and calculate rpms
  const queue = [[driver_id, driver_rpm]];
  visited.add(driver_id);
  rpms[driver_id] = driver_rpm;
  
  while (queue.length > 0) {
    const [current_id, current_rpm] = queue.shift();
    
    for (const next_id of adj[current_id]) {
      if (!visited.has(next_id)) {
        visited.add(next_id);
        // ratio is based on tooth counts of driving vs driven gear
        const ratio = gears[current_id] / gears[next_id];
        // direction reverses each time a gear drives another
        const next_rpm = -current_rpm * ratio;
        rpms[next_id] = next_rpm;
        queue.push([next_id, next_rpm]);
      }
    }
  }
  
  return rpms;
}