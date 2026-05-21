function map(head, fn) {
  if (head === null) return null;
  return new Node(fn(head.data), map(head.next, fn));
}