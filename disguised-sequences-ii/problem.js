/*
Let us define two sums v(n, p) and u(n, p):

v(n,p) = Σ(k=0 to n) (-1)^k × p × 4^(n-k) × C(2n-k, k)
u(n,p) = Σ(k=0 to n) (-1)^k × p × 4^(n-k) × C(2n-k+1, k)

where C(n, k) is the binomial coefficient.

Tasks:
1. Calculate v(n, p) and u(n, p) with two brute-force functions v1(n, p) and u1(n, p).
2. Try v1(n, p) and u1(n, p) for small values of n and p and guess the results.
3. Write vEff(n, p) and uEff(n, p) to efficiently calculate v and u.

Examples:
v1(12, 70) --> 1750
u1(13, 18) --> 252
*/

function binomial(n, k) {
  // TODO: implement
}

function v1(n, p) {
  // TODO: implement brute force
}

function u1(n, p) {
  // TODO: implement brute force
}

function vEff(n, p) {
  // TODO: implement efficient version
}

function uEff(n, p) {
  // TODO: implement efficient version
}