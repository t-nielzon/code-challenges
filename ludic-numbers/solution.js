function sumLudic(n) {
  const ludic = [1];
  let arr = [];
  
  // Generate starting array [2, 3, 4, 5, ...]
  for (let i = 2; i <= 300000; i++) {
    arr.push(i);
  }
  
  while (ludic.length < n && arr.length > 0) {
    const k = arr[0];
    ludic.push(k);
    
    // Remove every kth element (indices 0, k, 2k, 3k, ...)
    const newArr = [];
    for (let i = 0; i < arr.length; i++) {
      if (i % k !== 0) {
        newArr.push(arr[i]);
      }
    }
    arr = newArr;
  }
  
  // Sum the first n ludic numbers
  return ludic.slice(0, n).reduce((sum, num) => sum + num, 0);
}