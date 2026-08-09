/*
A vending machine is a machine that dispenses items such as snacks and beverages to customers 
automatically, after the customer inserts currency or credit into the machine.

Task: Recreate a vending machine as a class called VendingMachine with a vend method.
The vend method takes two arguments:
1. Selection code of the item (not case sensitive)
2. Amount of money the user inserts into the machine

Rules:
1. If money given < item cost return "Not enough money!"
2. If quantity is 0 return "Item Name: Out of stock!"
3. If correctly selected return "Vending Item Name with X.XX change."
4. If correctly selected with no change return "Vending Item Name"
5. If invalid item selected return "Invalid selection! : Money in vending machine = X.XX"
6. Update quantity on successful vend
7. Keep track of machine money
8. Change is always to 2 decimal places
*/

class VendingMachine {
  constructor(items, initialMoney) {
    // Implementation here
  }

  vend(selection, moneyInserted) {
    // Implementation here
  }
}