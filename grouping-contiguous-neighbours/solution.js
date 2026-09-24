function grouper(firstIndex, lastIndex, shouldMerge, groupValues, singleValue) {
  let i = firstIndex;
  
  while (i <= lastIndex) {
    // Check if current index should merge with the next
    if (i < lastIndex && shouldMerge(i)) {
      // Start of a group - collect consecutive mergeable indices
      let groupStart = i;
      while (i < lastIndex && shouldMerge(i)) {
        i++;
      }
      // Call groupValues for the entire consecutive range
      groupValues(groupStart, i);
      i++;
    } else {
      // Single, non-mergeable value
      singleValue(i);
      i++;
    }
  }
}