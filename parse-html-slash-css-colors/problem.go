/*
Parse HTML/CSS Colors

In this kata you parse RGB colors represented by strings. The formats are
primarily used in HTML and CSS. Your task is to implement a function which
takes a color as a string and returns the parsed color as a map.

Input:
- 6-digit hexadecimal - "#RRGGBB" e.g. "#012345", "#789abc", "#FFA077"
- 3-digit hexadecimal - "#RGB" e.g. "#012", "#aaa", "#F5A"
- Preset color name e.g. "red", "BLUE", "LimeGreen"

Examples:
  Parse("#80FFA0")   == Color{128, 255, 160}
  Parse("#3B7")      == Color{51, 187, 119}
  Parse("LimeGreen") == Color{50, 205, 50}
*/
package kata

// Color struct is defined as follows:
// type Color struct{ R, G, B byte }

// PresetColors is a predefined map with lower-case keys and "#RRGGBB" values.

func Parse(color string) Color {
}