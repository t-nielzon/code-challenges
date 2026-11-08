package main

/*
Phone Directory

John keeps a backup of his old personal phone book as a text file. On each line of the file
he can find the phone number (formatted as +X-abc-def-ghij where X stands for one or two digits),
the corresponding name between < and >, and the address.

Unfortunately everything is mixed, things are not always in the same order; parts of lines are
cluttered with non-alpha-numeric characters (except inside phone number and name).

Given the lines of his phone book and a phone number num, return a string:
"Phone => num, Name => name, Address => address"

If there are multiple people with the same phone number:
"Error => Too many people: num"

If the number is not found:
"Error => Not found: num"
*/

func phone(s string, num string) string {
	return ""
}