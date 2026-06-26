function bus(arr) {
  return arr.reduce((sum, [on, off]) => sum + on - off, 0);
}