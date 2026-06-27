function reverseNumber(n) {
  const len = n.length;

  // compare two equal-length numeric strings, return true if a >= b
  const geq = (a, b) => a >= b;

  // build palindrome by mirroring the first half onto the second half
  const mirror = (s) => {
    const arr = s.split('');
    for (let i = 0; i < Math.floor(len / 2); i++) {
      arr[len - 1 - i] = arr[i];
    }
    return arr.join('');
  };

  // a palindrome of the same digit-length reads the same reversed, so both
  // m and its reverse equal m itself; we need the smallest palindrome >= n
  let candidate = mirror(n);
  if (geq(candidate, n)) {
    return candidate;
  }

  // mirror was smaller than n: increment the left half (including middle
  // digit) and mirror again to get the next larger palindrome
  const half = Math.ceil(len / 2);
  let left = n.slice(0, half);

  // increment the left-half integer as a string (no overflow since n < 10^100)
  const incremented = (BigInt(left) + 1n).toString().padStart(half, '0');

  const arr = n.split('');
  for (let i = 0; i < half; i++) {
    arr[i] = incremented[i];
  }
  for (let i = 0; i < Math.floor(len / 2); i++) {
    arr[len - 1 - i] = arr[i];
  }
  return arr.join('');
}