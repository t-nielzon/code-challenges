/*
GA-DE-RY-PO-LU-KI cypher vol 3 - Missing key

## Introduction

The GADERYPOLUKI is a simple substitution cypher used in scouting to encrypt messages.
The encryption is based on short, easy to remember key. The key is written as paired
letters, which are in the cipher simple replacement.

The most frequently used key is "GA-DE-RY-PO-LU-KI".

 g => a
 a => g
 d => e
 e => d
  etc.

The letters, which are not on the list of substitutes, stay in the encrypted text
without changes.

## Task

Our scouts had a party yesterday, and they had too much milk and cookies. As a result,
all of them forgot the key. Your task is to help scouts to find the key that they used
for encryption. Fortunately, they have some messages that are already encoded.

## Input

The function accepts two arrays.

The `messages` string array consists of lowercase characters and whitespace characters.
The strings on the `messages` array are scout's messages before encryption.

The `secrets` string array consists of lowercase characters and whitespace characters.
The strings on the `secrets` array are scout's messages after encryption.

## Output

The returned string should consist of lowercase characters only. The pairs of substitutions
should be ordered by the first letter of substitution. The letters in each pair should be
in alphabetical order.

## Example

var messages = []string{"dance on the table", "hide my beers", "scouts rocks"}
var secrets = []string{"egncd pn thd tgbud", "hked mr bddys", "scplts ypcis"}
FindTheKey(messages, secrets)   //=> agdeikluopry
*/
package kata

func FindTheKey(messages, secrets []string) string {
}