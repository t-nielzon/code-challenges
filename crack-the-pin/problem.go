package main

/*
Crack the PIN

Given is a md5 hash of a five digits long PIN. It is given as string.
Md5 is a function to hash your password:
"password123" ===> "482c811da5d5b4bc6d497ffa98491e38"

Why is this useful?
Hash functions like md5 can create a hash from string in a short time and it is
impossible to find out the password, if you only got the hash. The only way is
cracking it, means try every combination, hash it and compare it with the hash
you want to crack.

Your task is to return the cracked PIN as string.

The PIN is 5 digits long (00000 to 99999).
*/

func CrackPin(hash string) string {
	return ""
}