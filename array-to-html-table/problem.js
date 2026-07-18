/*
Overview:
Given a 2D array (list), generate an HTML table representing the data from this array.

You need to write a function called toTable, that takes three arguments:
- data: a 2D array (list)
- headers: an optional boolean value. If true, the first row of the array is considered a header. Defaults to false.
- index: an optional boolean value. If true, the first column in the table should contain 1-based indices of the corresponding row. If headers is true, this column should have an empty header. Defaults to false.

Returns a string containing HTML tags representing the table (minified, no whitespace).

Details:
- If headers is true, use <thead> with <th> elements for the first row
- If headers is false, the table starts with <tbody>
- If index is true, add a first column with 1-based indices
- null values should be converted to empty strings
- Other values should be converted to their string representation

Example:
toTable([["lorem", "ipsum"], ["dolor", "sit amet"]], true, true)
returns: "<table><thead><tr><th></th><th>lorem</th><th>ipsum</th></tr></thead><tbody><tr><td>1</td><td>dolor</td><td>sit amet</td></tr></tbody></table>"
*/

function toTable(data, headers = false, index = false) {
  
}