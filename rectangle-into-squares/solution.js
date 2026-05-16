function sqInRect(lng, wdth) {
  if (lng === wdth) return null;
  const result = [];
  while (lng > 0 && wdth > 0) {
    if (lng > wdth) {
      result.push(wdth);
      lng -= wdth;
    } else {
      result.push(lng);
      wdth -= lng;
    }
  }
  return result;
}