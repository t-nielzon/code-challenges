class VendingMachine {
  constructor(items, money) {
    this.items = items;
    this.money = money;
  }

  vend(code, moneyInserted) {
    const itemCode = code.toUpperCase();
    const item = this.items.find(i => i.code.toUpperCase() === itemCode);

    if (!item) {
      return `Invalid selection! : Money in vending machine = ${this.money.toFixed(2)}`;
    }

    if (item.quantity === 0) {
      return `${item.name}: Out of stock!`;
    }

    if (moneyInserted < item.price) {
      return "Not enough money!";
    }

    const change = moneyInserted - item.price;
    item.quantity--;
    this.money += moneyInserted;

    if (change > 0) {
      return `Vending ${item.name} with ${change.toFixed(2)} change.`;
    } else {
      return `Vending ${item.name}`;
    }
  }
}