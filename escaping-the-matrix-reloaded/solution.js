// solution.js
class Neo extends Error {
  constructor() {
    super('Escaped!');
    this.name = 'Neo';
  }
}

function escapeTheMatrix() {
  throw new Neo();
}