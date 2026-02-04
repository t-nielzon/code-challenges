/*
⛏️ **The Pickaxe Collector**

As a collector, you aim to improve your pickaxes by trading your current ones for a better-performing pickaxe, while staying within your budget.

### Rules 

- Buy the strongest pickaxe you can afford, selling from your inventory if needed.
- Only buy if it is strictly stronger than your current best.
- Your current best is the highest-damage pickaxe you own; if you own none, all store pickaxes count as better.
- Each pickaxe has a unique damage value.  
- You may hold only one pickaxe per material.  
- When selling, start from the lowest-damage pickaxe and continue upward until you can afford the best option.  
- Your budget cannot go below zero.  
- Pickaxes with `sell_price = 0` cannot be sold — these are worthless relics.  
- Pickaxes unknown to the store cannot be sold — store owners show no interest in them. 
- Known materials are always in stock, and the store has unlimited funds.  
- For all pickaxes, `buy_price > sell_price`.

### Task

Decide which pickaxe to buy and which to sell, if any, to obtain the **strongest pickaxe you can afford** under the given rules.

### Inputs

- **materials:** A dictionary mapping `material_name` → `damage`. (all materials from store and inventory will be available)  
- **store:** A dictionary mapping `material_name` → `(buy_price, sell_price)`.  
- **inventory:** A set of material names you currently own (each at most once).  
- **budget:** A nonnegative integer representing your available funds.

### Output

- Either:
  - A tuple `(bought_name, sold_set, remaining_budget)` where:
    - `bought_name` is the material name of the pickaxe you buy.  
    - `sold_set` is the set of material names you sell to afford it (possibly empty).  
    - `remaining_budget` is the budget remaining after trading pickaxes.  
- Or:
  - `None` if you did not buy a pickaxe.
*/

function pickaxeCollector(materials, store, inventory, budget) {
  // Your solution here
}

module.exports = { pickaxeCollector };