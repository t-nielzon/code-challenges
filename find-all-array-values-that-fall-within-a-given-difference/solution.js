class GroupByDifference {
  constructor(numbers) {
    this.numbers = numbers;
  }

  find(difference) {
    const sorted = [...this.numbers].sort((a, b) => a - b);
    const keep = new Array(sorted.length).fill(false);

    for (let i = 0; i < sorted.length; i++) {
      for (let j = 0; j < sorted.length; j++) {
        if (i === j) continue;
        if (Math.abs(sorted[i] - sorted[j]) <= difference) {
          keep[i] = true;
          keep[j] = true;
        }
      }
    }

    return sorted.filter((_, i) => keep[i]);
  }
}