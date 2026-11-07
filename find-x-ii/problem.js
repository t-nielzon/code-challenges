/*
  Optimize the recursive findX function for large numbers.
  
  function findX(n) {
    if (n == 0) return 0;
    let x = 0;
    for (let i=1; i<=n; ++i)
      x += findX(i-1) + 3*i;
    return x;
  }
  
  Test cases:
  findX(2) //=> 12
  findX(3) //=> 33
  findX(5) //=> 171
  
  Return result % (10^9+7)
  Input: 1 <= n <= 10^6
*/

function findX(n) {
  // optimize this
}