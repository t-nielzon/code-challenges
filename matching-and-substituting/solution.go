Here we will work with strings like the string data above and not with files.

The function `change(s, prog, version)` given:

`s=data, prog="Ladder" , version="1.1"` will return:

`"Program: Ladder Author: g964 Phone: +1-503-555-0090 Date: 2019-01-01 Version: 1.1"`

Rules:

- The date should always be `"2019-01-01"`.
- The author should always be `"g964"`.
- Replace the current `"Program Title"` with the prog argument supplied to your function.
  Also remove `"Title"`.
- Remove the lines containing `"Corporation"` and `"Level"` completely.
- Phone numbers and versions must be in valid formats.
- A valid version in the input is one or more digits followed by a dot, followed by one or more digits.
- A valid input phone format is +1-xxx-xxx-xxxx, where each x is a digit.
- If the phone or version format is not valid, return `"ERROR: VERSION or PHONE"`.
- If the version format is valid and the version is anything other than `2.0`, replace it with the version parameter.
  If it's `2.0`, don't modify it.
- If the phone number is valid, replace it with `"+1-503-555-0090"`.
*/

package kata

func Change(s, prog, version string) string {
	return ""
}