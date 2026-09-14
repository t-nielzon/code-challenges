class Neo extends Error {
  constructor(message) {
    super(message);
    this.name = 'Neo';
  }
}

function matrix() {
  throw new Neo();
}