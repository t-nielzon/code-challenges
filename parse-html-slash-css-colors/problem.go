/*
Parse HTML/CSS Colors

In this kata you parse RGB colors represented by strings. The formats are
primarily used in HTML and CSS. Your task is to implement a function which
takes a color as a string and returns the parsed color as a Color struct.

Input:
The input string represents one of the following:

1. 6-digit hexadecimal - "#RRGGBB"
   e.g. "#012345", "#789abc", "#FFA077"
   Each pair of digits represents a value of the channel in hexadecimal: 00 to FF

2. 3-digit hexadecimal - "#RGB"
   e.g. "#012", "#aaa", "#F5A"
   Each digit represents a value 0 to F which translates to 2-digit hexadecimal:
   0->00, 1->11, 2->22, and so on.

3. Preset color name
   e.g. "red", "BLUE", "LimeGreen"
   You have to use the predefined map PresetColors. The keys are the names of
   preset colors in lower-case and the values are the corresponding colors in
   6-digit hexadecimal (same as "#RRGGBB").

Examples:
   Parse("#80FFA0")   == Color{128, 255, 160}
   Parse("#3B7")      == Color{51, 187, 119}
   Parse("LimeGreen") == Color{50, 205, 50}

Color struct is defined as follows:
   type Color struct{ R, G, B byte }
*/

package kata

type Color struct{ R, G, B byte }

func Parse(s string) Color {
}