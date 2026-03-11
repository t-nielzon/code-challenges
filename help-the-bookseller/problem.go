/*
A bookseller has lots of books classified in 26 categories labeled A, B, C, ..., Z.
Each book has a code of at least 3 characters. The 1st character of a code is a
capital letter which defines the book category.

In the bookseller's stocklist each code is followed by a space and by a positive
integer, which indicates the quantity of books of this code in stock.

You will receive the bookseller's stocklist and a list of categories. Your task is
to find the total number of books in the bookseller's stocklist, with the category
codes in the list of categories.

Return the result as a string. If any of the input lists is empty, return an empty string.

Example:
  stocklist: "ABART 20", "CDXEF 50", "BKWRK 25", "BTSQZ 89", "DRTYM 60"
  categories: "A", "B", "C", "W"
  result: "(A : 20) - (B : 114) - (C : 50) - (W : 0)"
*/
package kata

func StockList(listOfArt []string, listOfCat []string) string {
}