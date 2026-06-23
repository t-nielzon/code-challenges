function eqAll(list) {
  let first;
  let hasFirst = false;
  for (const item of list) {
    if (!hasFirst) {
      first = item;
      hasFirst = true;
    } else if (item !== first) {
      return false;
    }
  }
  return true;
}