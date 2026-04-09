function loop_size(node) {
  let slow = node;
  let fast = node;

  // detect cycle using floyd's tortoise and hare
  do {
    slow = slow.next;
    fast = fast.next.next;
  } while (slow !== fast);

  // count the loop length from the meeting point
  let count = 1;
  slow = slow.next;
  while (slow !== fast) {
    slow = slow.next;
    count++;
  }

  return count;
}

module.exports = loop_size;