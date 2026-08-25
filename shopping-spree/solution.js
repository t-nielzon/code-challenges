function shoppingSpree(p, shop) {
  let maxUnlocked = 0;
  let sum = 0;
  let largest1 = 0, largest2 = 0;
  
  for (let end = 0; end < shop.length; end++) {
    sum += shop[end];
    
    // Update the two largest items seen so far
    if (shop[end] >= largest1) {
      largest2 = largest1;
      largest1 = shop[end];
    } else if (shop[end] > largest2) {
      largest2 = shop[end];
    }
    
    // Try using 0 coupons
    if (sum <= p) {
      maxUnlocked = end + 1;
      continue;
    }
    
    // Try using 1 coupon (on the most expensive item)
    if (sum - largest1 <= p) {
      maxUnlocked = end + 1;
      continue;
    }
    
    // Try using 2 coupons (on the two most expensive items)
    if (sum - largest1 - largest2 <= p) {
      maxUnlocked = end + 1;
    }
  }
  
  return maxUnlocked;
}