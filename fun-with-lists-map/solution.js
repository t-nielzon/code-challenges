function map(head, mapper) {
  if (head === null) return null;
  return new Node(mapper(head.data), map(head.next, mapper));
}