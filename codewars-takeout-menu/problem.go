/*
## CodeWars Takeout Menu

### Customer Dialogue:
- The customer will always open the conversation with `place an order`
- The customer may ask for a `price {item}`
- The customer confirms they want an item by saying `order {item}`
- The customer may ask for an item that is not on the menu
- The order ends when the customer says `that is all`

### Expected Responses:
- Respond with `ok` for
  - `place an order`
  - `order {item}` - (`unavailable` if not on the menu)
- Respond with the price as a string for
  - `price {item}` - (`unavailable` if not on the menu)
  - `that is all`
- Respond to `goodbye` with `goodbye`

### When to stop
You should continue to take orders until you receive on the `done` channel.

### Example Phone Call
place an order
    ok
price chicken bento box
    12.99
price spring roll
    1.99
order spring roll
    ok
order spring roll
    ok
order chicken bento box
    ok
price pizza
    unavailable
that is all
    16.97
goodbye
    goodbye
*/
package kata

func Solution(hear <-chan string, speak chan<- string, done <-chan bool) {
}