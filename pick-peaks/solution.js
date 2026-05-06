function pickPeaks(arr) {
  const pos = [];
  const peaks = [];
  let lastIncreasePos = -1;

  for (let i = 1; i < arr.length; i++) {
    if (arr[i] > arr[i - 1]) {
      lastIncreasePos = i;
    } else if (arr[i] < arr[i - 1] && lastIncreasePos !== -1) {
      pos.push(lastIncreasePos);
      peaks.push(arr[lastIncreasePos]);
      lastIncreasePos = -1;
    }
  }

  return { pos, peaks };
}