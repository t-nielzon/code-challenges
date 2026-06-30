class GroupByDifference {
  constructor(numbers) {
    this.numbers = numbers;
  }

  find(difference) {
    const sorted = [...this.numbers].sort((a, b) => a - b);
    const result = [];
    let group = [];

    for (let i = 0; i < sorted.length; i++) {
      if (group.length === 0 || sorted[i] - sorted[i - 1] <= difference) {
        group.push(sorted[i]);
      } else {
        if (group.length > 1) result.push(...group);
        group = [sorted[i]];
      }
    }
    if (group.length > 1) result.push(...group);

    return result;
  }
}