/*
John keeps a backup of his old personal phone book as a text file. On each line of
the file he can find the phone number (formated as `+X-abc-def-ghij` where X stands
for one or two digits), the corresponding name between `<` and `>` and the address.

Unfortunately everything is mixed, things are not always in the same order; parts of
lines are cluttered with non-alpha-numeric characters (except inside phone number and
name).

Examples of John's phone book lines:

  "/+1-541-754-3010 156 Alphand_St. <J Steeve>\n"
  " 133, Green, Rd. <E Kustur> NY-56423 ;+1-541-914-3010!\n"
  "<Anastasia> +48-421-674-8974 Via Quirinal Roma\n"

Given the lines of his phone book and a phone number `num`, return a string for this
number: "Phone => num, Name => name, Address => adress"

Examples:

  s = "/+1-541-754-3010 156 Alphand_St. <J Steeve>\n 133, Green, Rd. <E Kustur> NY-56423 ;+1-541-914-3010!\n"
  phone(s, "1-541-754-3010") -> "Phone => 1-541-754-3010, Name => J Steeve, Address => 156 Alphand St."

- If there are many people for a phone number `num`, return: "Error => Too many people: num"
- If the number `num` is not in the phone book, return: "Error => Not found: num"
*/

function phone(strng, num) {
  // your code here
}