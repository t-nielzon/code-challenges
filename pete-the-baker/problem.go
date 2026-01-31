/*
Pete, the baker

Pete likes to bake some cakes. He has some recipes and ingredients.
Unfortunately he is not good in maths. Can you help him to find out,
how many cakes he could bake considering his recipes?

Write a function Cakes(), which takes the recipe (map) and the available
ingredients (also a map) and returns the maximum number of cakes Pete
can bake (integer). For simplicity there are no units for the amounts
(e.g. 1 lb of flour or 200 g of sugar are simply 1 or 200). Ingredients
that are not present in the objects, can be considered as 0.

Examples:
// must return 2
Cakes(map[string]int{"flour": 500, "sugar": 200, "eggs": 1},
      map[string]int{"flour": 1200, "sugar": 1200, "eggs": 5, "milk": 200})

// must return 0
Cakes(map[string]int{"apples": 3, "flour": 300, "sugar": 150, "milk": 100, "oil": 100},
      map[string]int{"sugar": 500, "flour": 2000, "milk": 2000})
*/
package kata

func Cakes(recipe, available map[string]int) int {
}