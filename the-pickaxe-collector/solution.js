function pickaxeCollector(materials, store, inventory, budget) {
  // Find current best damage in inventory
  let currentBestDamage = -1;
  for (const material of inventory) {
    if (materials[material] > currentBestDamage) {
      currentBestDamage = materials[material];
    }
  }

  // Get sellable items from inventory (in store and sell_price > 0)
  // Sorted by damage ascending (sell weakest first)
  const sellable = [];
  for (const material of inventory) {
    if (store[material] && store[material][1] > 0) {
      sellable.push({
        name: material,
        damage: materials[material],
        sellPrice: store[material][1]
      });
    }
  }
  sellable.sort((a, b) => a.damage - b.damage);

  // Get store items sorted by damage descending (try strongest first)
  const storeItems = [];
  for (const material in store) {
    storeItems.push({
      name: material,
      damage: materials[material],
      buyPrice: store[material][0]
    });
  }
  storeItems.sort((a, b) => b.damage - a.damage);

  // Try to buy the strongest pickaxe we can afford
  for (const item of storeItems) {
    // Only consider if strictly stronger than current best
    if (item.damage <= currentBestDamage) {
      continue;
    }

    // Skip if we already own this material
    if (inventory.has(item.name)) {
      continue;
    }

    // Try to afford this item by selling from weakest up
    let funds = budget;
    const toSell = new Set();

    // Sell items starting from lowest damage until we can afford or run out
    for (const sellItem of sellable) {
      if (funds >= item.buyPrice) {
        break;
      }
      // Don't sell something stronger than what we're buying
      if (sellItem.damage >= item.damage) {
        continue;
      }
      funds += sellItem.sellPrice;
      toSell.add(sellItem.name);
    }

    if (funds >= item.buyPrice) {
      const remaining = funds - item.buyPrice;
      return [item.name, toSell, remaining];
    }
  }

  return null;
}

module.exports = { pickaxeCollector };