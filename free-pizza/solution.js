function pizzaRewards(customers, minOrders, minPrice) {
  const eligible = new Set();
  for (const [name, orders] of Object.entries(customers)) {
    if (orders.filter(o => o >= minPrice).length >= minOrders) {
      eligible.add(name);
    }
  }
  return eligible;
}

module.exports = { pizzaRewards };