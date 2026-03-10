/*
Second Variation on Caesar Cipher

In this country soldiers are poor but they need a certain level of secrecy
for their communications so, though they do not know Caesar cypher, they
reinvent it in the following way.

The action of a Caesar cipher is to replace each plaintext letter with a
different one a fixed number of places up or down the alphabet. This
displacement is called the "shift" or "rotate". The shift is cyclic.

They code only letters a-z and A-Z. Other characters are kept as is.

They change the "rotate" each new message. This "rotate" is a prefix for
their message once coded. The prefix is built of 2 letters, the second
shifted from the first by the rotate, the first being the downcased first
letter of the uncoded message.

To lessen risk they cut the coded message and prefix into five pieces.
If possible evenly split; if not, parts 1-4 are longer and part 5 shorter.
Choose the option where part 5 has the longest length. If the last part is
empty, don't include it in the result.

Example with shift = 1:
  message: "I should have known that you would have a perfect answer for me!!!"
  code: ["ijJ tipvme ibw", "f lopxo uibu z", "pv xpvme ibwf ", "b qfsgfdu botx", "fs gps nf!!!"]
*/
package kata

func Encode(s string, shift int) []string {
	return nil
}

func Decode(arr []string) string {
	return ""
}