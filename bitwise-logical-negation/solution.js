function negation(n) {
  return 1+((n|~n+1)>>31)
}