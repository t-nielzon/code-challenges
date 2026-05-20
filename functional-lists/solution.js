function EmptyList() {
}

EmptyList.prototype.toString = function () {
  return "()";
};

EmptyList.prototype.isEmpty = function () {
  return true;
};

EmptyList.prototype.length = function () {
  return 0;
};

EmptyList.prototype.push = function (value) {
  return new ListNode(value, this);
};

EmptyList.prototype.remove = function (value) {
  return this;
};

EmptyList.prototype.append = function (other) {
  return other;
};

function ListNode(value, next) {
  this.value = value;
  this.next = next;
}

ListNode.prototype.toString = function () {
  var parts = [];
  var node = this;
  while (!node.isEmpty()) {
    parts.push(String(node.value));
    node = node.next;
  }
  return "(" + parts.join(" ") + ")";
};

ListNode.prototype.isEmpty = function () {
  return false;
};

ListNode.prototype.length = function () {
  var n = 0;
  var node = this;
  while (!node.isEmpty()) {
    n++;
    node = node.next;
  }
  return n;
};

ListNode.prototype.push = function (value) {
  return new ListNode(value, this);
};

ListNode.prototype.remove = function (value) {
  if (this.value === value) {
    return this.next.remove(value);
  }
  var newNext = this.next.remove(value);
  if (newNext === this.next) {
    return this;
  }
  return new ListNode(this.value, newNext);
};

ListNode.prototype.append = function (other) {
  return new ListNode(this.value, this.next.append(other));
};

ListNode.prototype.head = function () {
  return this.value;
};

ListNode.prototype.tail = function () {
  return this.next;
};