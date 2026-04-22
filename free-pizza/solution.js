function freePizza(customers, minOrders, minValue) {
  const result = new Set();
  for (const [name, orders] of Object.entries(customers)) {
    if (orders.filter(v => v >= minValue).length >= minOrders) {
      result.add(name);
    }
  }
  return result;
}