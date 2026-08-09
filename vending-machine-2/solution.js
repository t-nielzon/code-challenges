class VendingMachine {
  constructor(items, initialMoney) {
    this.items = items;
    this.money = initialMoney;
  }

  vend(selection, moneyInserted) {
    // Find the item (case insensitive)
    const item = this.items.find(i => i.code.toLowerCase() === selection.toLowerCase());
    
    // Check if item exists
    if (!item) {
      return `Invalid selection! : Money in vending machine = ${this.money.toFixed(2)}`;
    }
    
    // Check if out of stock
    if (item.quantity === 0) {
      return `${item.name}: Out of stock!`;
    }
    
    // Check if enough money
    if (moneyInserted < item.price) {
      return "Not enough money!";
    }
    
    // Calculate change and update state
    const change = moneyInserted - item.price;
    this.money += moneyInserted;
    item.quantity--;
    
    // Return message with or without change
    if (change > 0) {
      return `Vending ${item.name} with ${change.toFixed(2)} change.`;
    } else {
      return `Vending ${item.name}`;
    }
  }
}