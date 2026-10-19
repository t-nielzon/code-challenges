function findUniq(arr) {
  // Determine the common number by checking first 3 elements
  let common;
  if (arr[0] === arr[1]) {
    common = arr[0];
  } else if (arr[0] === arr[2]) {
    common = arr[0];
  } else {
    common = arr[1];
  }
  
  // Find and return the unique number
  for (let num of arr) {
    if (num !== common) {
      return num;
    }
  }
}