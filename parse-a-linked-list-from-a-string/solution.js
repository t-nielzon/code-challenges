class Node {
  constructor(data, next = null) {
    this.data = data;
    this.next = next;
  }
}

function parse(string) {
  const parts = string.split(' -> ');
  let head = null;
  for (let i = parts.length - 1; i >= 0; i--) {
    if (parts[i] === 'null') continue;
    head = new Node(parseInt(parts[i], 10), head);
  }
  return head;
}

module.exports = { Node, parse };