package main

/*
You are given a small extract of a catalog in XML format with products.
Each product has a name, price (prx), and quantity (qty).
Given an article name, return all matching products formatted as:
  "{name} > prx: ${price} qty: {qty}"
joined by newlines with a trailing newline.

If no products match, return "Nothing".

Example:
  s = "<prod><name>drill</name><prx>99</prx><qty>5</qty></prod>..."
  Catalog(s, "saw") returns "table saw > prx: $1099.99 qty: 5\nsaw > prx: $9 qty: 10\n"

Notes:
- Match products where the name contains the article string
- Preserve the order from the catalog
- Return "Nothing" if no matches found
*/
func Catalog(s, article string) string {
}