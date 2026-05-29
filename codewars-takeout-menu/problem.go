/*
## CodeWars Takeout Menu

Your task at the CodeWars restaurant is to answer phone calls and take customer
orders, telling them the final cost at the end of each phone call. The catch
with this kata though is that it uses channels to communicate with your
Solution rather than the typical single function call and return.

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
Continue to take orders until you receive on the `done` channel.

There is a preloaded function GetMenu() that returns a map[string]float32
containing the full menu + pricing.
*/

package kata

func TakeOrders(hear <-chan string, say chan<- string, done <-chan bool) {
}