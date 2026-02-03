function pickaxeCollector(materials, store, inventory, budget) {
  // Find current best damage from inventory
  let currentBestDamage = 0;
  for (const material of inventory) {
    if (materials[material] > currentBestDamage) {
      currentBestDamage = materials[material];
    }
  }
  
  // Get sellable items from inventory (in store and sell_price > 0), sorted by damage ascending
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
  
  // Get store items sorted by damage descending (strongest first)
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
    // Only buy if strictly stronger than current best
    if (item.damage <= currentBestDamage) {
      continue;
    }
    
    // Skip if we already own this material
    if (inventory.has(item.name)) {
      continue;
    }
    
    // Check if we can afford it, possibly by selling
    let currentBudget = budget;
    const toSell = new Set();
    
    // Try selling from lowest damage upward until we can afford
    for (const sellableItem of sellable) {
      if (currentBudget >= item.buyPrice) {
        break;
      }
      currentBudget += sellableItem.sellPrice;
      toSell.add(sellableItem.name);
    }
    
    // If we can afford it now, buy it
    if (currentBudget >= item.buyPrice) {
      const remainingBudget = currentBudget - item.buyPrice;
      return [item.name, toSell, remainingBudget];
    }
  }
  
  // Couldn't buy anything better
  return null;
}