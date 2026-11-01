function simunq(str) {
  // Replace disallowed characters with dash
  let simplified = str.replace(/[^a-zA-Z0-9\s-]/g, '-');
  
  // Check if we've returned this simplified string before
  if (simunq.returned.has(simplified)) {
    simunq.returned.set(simplified, simunq.returned.get(simplified) + 1);
    const count = simunq.returned.get(simplified);
    const result = simplified + '-' + count;
    simunq.returned.set(result, 1);
    return result;
  } else {
    simunq.returned.set(simplified, 1);
    return simplified;
  }
}

simunq.returned = new Map();

simunq.reset = function() {
  simunq.returned.clear();
};