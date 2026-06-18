/*
 * Formatting simple HTML (5 kyu)
 *
 * Task:
 *   Format a source in a simple HTML-dialect. A source consists of tags and text.
 *   It is not necessarily valid HTML, or a complete document (it may be a snippet).
 *
 * Tags and text:
 *   Tags are either matching opening/closing tags (<tag>content</tag>), or
 *   self-closed (<tag />). All tags need to be on their own line. There are no
 *   inline tags. Content between opening and closing tags should be indented.
 *   No whitespace (other than indent/newline) should be before or after tags.
 *
 *   All (consecutive) text needs to be on its own line. Text may have spurious
 *   whitespace; this needs to be collapsed to single spaces. (Do not reformat
 *   whitespace inside tags.) Text should not begin or end with whitespace
 *   (other than indent/newline).
 *
 *   Exception: the <br /> tag, when not after a tag, should be after its text
 *   without an intervening newline. It should be followed by a newline.
 *
 * Details:
 *   - Newlines must be \n
 *   - Indents must be two spaces per level
 *   - Whitespace in text must be single spaces
 *   - Reformatted source must end with a newline
 *   - Repeated application should not change the output (idempotent)
 *   - All input is valid
 *   - There will be no whitespace in tags until after the tag name, and none
 *     directly before the closing > (tag> or />)
 */

function prettyfy(source) {
  // your code here
}