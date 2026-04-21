function solution(arr1, arr2) {
  return arr1.reduce((sum, v, i) => sum + (v - arr2[i]) ** 2, 0) / arr1.length;
}