function zipWith(fn, list1, list2) {
  if (list1 === null || list2 === null) {
    return null;
  }
  
  return new Node(
    fn(list1.value, list2.value),
    zipWith(fn, list1.next, list2.next)
  );
}