function queueTime(customers, n) {
  const tills = new Array(n).fill(0);

  for (const time of customers) {
    // assign each customer to the till that becomes free earliest
    const earliest = tills.indexOf(Math.min(...tills));
    tills[earliest] += time;
  }

  return Math.max(...tills);
}