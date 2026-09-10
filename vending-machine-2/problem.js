/*
# Vending Machine (6 kyu)

A vending machine is a machine that dispenses items such as snacks and beverages to customers automatically,
after the customer inserts currency or credit into the machine.

## Task

Create a VendingMachine class with a vend method that takes a selection code and money amount.

## Rules

1. If money < item cost: return "Not enough money!"
2. If quantity is 0: return "Item Name: Out of stock!"
3. If valid with change: return "Vending Item Name with X.XX change."
4. If valid with no change: return "Vending Item Name"
5. If invalid selection: return "Invalid selection! : Money in vending machine = X.XX"
6. Update quantity on successful vend
7. Track machine money
8. Change formatted to 2 decimal places
*/

class VendingMachine {
  constructor(items, money) {
    // Initialize with items and initial machine money
  }

  vend(code, moneyInserted) {
    // Process vending request
  }
}