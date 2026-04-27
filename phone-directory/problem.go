/*
John keeps a backup of his old personal phone book as a text file. On each line of the file
he can find the phone number (formated as +X-abc-def-ghij where X stands for one or two digits),
the corresponding name between < and > and the address.

Unfortunately everything is mixed, things are not always in the same order; parts of lines are
cluttered with non-alpha-numeric characters (except inside phone number and name).

Given the lines of his phone book and a phone number num returns a string for this number:
"Phone => num, Name => name, Address => adress"

If there are many people for a phone number num, return: "Error => Too many people: num"
If the number num is not in the phone book, return: "Error => Not found: num"
*/

package kata

func phone(strng string, num string) string {
}