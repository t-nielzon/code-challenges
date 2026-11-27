function queueTime(customers, n) {
  const tills = new Array(n).fill(0);
  
  for (const time of customers) {
    const minIndex = tills.indexOf(Math.min(...tills));
    tills[minIndex] += time;
  }
  
  return Math.max(...tills);
}