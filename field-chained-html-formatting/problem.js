/*
 * Field Chained HTML Formatting
 *
 * We want to create an object with methods for various HTML elements: div, p, span and h1.
 * These methods will wrap the passed-in string in the tag associated with each.
 *
 * Format.div("foo"); // returns "<div>foo</div>"
 * Format.p("bar"); // returns "<p>bar</p>"
 *
 * We also want to be able to add additional formatting by chaining our methods together.
 *
 * Format.div.h1("FooBar"); // "<div><h1>FooBar</h1></div>"
 * Format.div.p.span("FizBuz"); // "<div><p><span>FizBuz</span></p></div>"
 *
 * Multiple arguments should be concatenated and wrapped in the correct HTML formatting.
 *
 * Format.div.h1("Foo", "Bar"); // "<div><h1>FooBar</h1></div>"
 *
 * We should be able to store the created methods and reuse them.
 *
 * var wrapInDiv = Format.div;
 * wrapInDiv("Foo");   // "<div>Foo</div>"
 * wrapInDiv.p("Bar"); // "<div><p>Bar</p></div>"
 *
 * And finally, we should be able to nest calls.
 *
 * Format.div(Format.h1("Title"), Format.p("paragraph"))
 * // "<div><h1>Title</h1><p>paragraph</p></div>"
 */

var Format = // implement Format object