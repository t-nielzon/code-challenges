function queueTime(customers, n) {
  const tills = Array(n).fill(0);
  
  for (const time of customers) {
    let minIndex = 0;
    for (let i = 1; i < n; i++) {
      if (tills[i] < tills[minIndex]) {
        minIndex = i;
      }
    }
    tills[minIndex] += time;
  }
  
  return Math.max(...tills);
}