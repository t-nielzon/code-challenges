function findEvenIndex(arr) {
  const totalSum = arr.reduce((sum, num) => sum + num, 0);
  let leftSum = 0;
  
  for (let i = 0; i < arr.length; i++) {
    const rightSum = totalSum - leftSum - arr[i];
    
    if (leftSum === rightSum) {
      return i;
    }
    
    leftSum += arr[i];
  }
  
  return -1;
}