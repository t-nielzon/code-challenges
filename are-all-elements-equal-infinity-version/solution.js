function eqAll(list) {
  let first = null;
  let isFirst = true;
  
  for (const element of list) {
    if (isFirst) {
      first = element;
      isFirst = false;
    } else if (element !== first) {
      return false;
    }
  }
  
  return true;
}