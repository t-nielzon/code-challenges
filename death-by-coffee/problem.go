/*
# Background

I drink too much coffee. Eventually it will probably kill me.

*Or will it..?*

Anyway, there's no way to know.

*Or is there...?*

# The Discovery of the Formula

I proudly announce my discovery of a formula for measuring the life-span of coffee drinkers!

For

* h is a health number assigned to each person (8 digit date of birth YYYYMMDD)
* CAFE is a cup of *regular* coffee
* DECAF is a cup of *decaffeinated* coffee

To determine the life-time coffee limits:

* Drink cups of coffee (i.e. add to h) until any part of the health number includes DEAD
* If the test subject can survive drinking five thousand cups without being DEAD then coffee has no ill effect on them

# Kata Task

Given the test subject's date of birth, return an array describing their life-time coffee limits

[ regular limit , decaffeinated limit ]

## Notes

* The limits are 0 if the subject is unaffected as described above
* At least 1 cup must be consumed (Just thinking about coffee cannot kill you!)

# Examples

* John was born 19/Jan/1950 so h=19500119. His coffee limits are [2645, 1162] which is only about 1 cup per week.
* Susan (11/Dec/1965) is unaffected by decaffeinated coffee, but regular coffee is very bad for her [111, 0].
* Elizabeth (28/Nov/1964) is allergic to decaffeinated coffee. Dead after only 11 cups [0, 11].
* Peter (4/Sep/1965) can drink as much coffee as he likes [0, 0].

Hint: https://en.wikipedia.org/wiki/Hexadecimal
*/
package kata

func CoffeeLimits(birthDate int) [2]int {
}