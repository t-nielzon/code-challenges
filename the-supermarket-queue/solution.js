function queueTime(customers, n) {
  const tills = Array(n).fill(0);
  
  for (const customer of customers) {
    const minIndex = tills.indexOf(Math.min(...tills));
    tills[minIndex] += customer;
  }
  
  return Math.max(...tills);
}