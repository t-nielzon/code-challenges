/*
 * Extract the contents of a given tag from the HTML document
 * 
 * Input:
 * - Parameter 1: An HTML markup string
 * - Parameter 2: A string representing an HTML tag name (e.g., 'h2')
 * 
 * Output:
 * An array of strings which represents the contents between the opening and closing tag.
 * Contents are trimmed of whitespace and newlines.
 * Returns an empty array if the tag is not found.
 * 
 * Assumptions:
 * - No angled brackets inside quoted attributes
 * - Tags specified will never contain nested tags, only inner text
 * - Contents may be on multiple lines (whitespace should be trimmed)
 * 
 * Examples:
 * getTagContent(htmlString, 'h1') --> ["Nature's Wonders"]
 * getTagContent(htmlString, 'h2') --> ["Birds","Butterflies"]
 * getTagContent(htmlString, 'p') --> ["In this article we discuss animals.", ...]
 * getTagContent(htmlString, 'h3') --> []
 */

function getTagContent(html, tag) {
  // TODO: implement
}