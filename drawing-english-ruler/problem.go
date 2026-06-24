// ## Draw an English Ruler (ASCII)
//
// A standard English ruler is marked in inches. For each whole inch, the
// ruler shows a major tick with a numeric label. Each time the interval
// size is halved, the tick length decreases by 1. The length of the major
// tick `t` defines the scale of the ruler.
//
// Write a function that generates the ASCII drawing of a ruler.
//
//   * `t` - major tick length
//   * `n` - number of inches on the ruler
//
// ## Examples
//
//   t=2 (n=4)      t=3 (n=3)
//   -- 0           --- 0
//   -              -
//   -- 1           --
//   -              -
//   -- 2           --- 1
//   -              -
//   -- 3           --
//   -              -
//   -- 4           --- 2
//                  -
//                  --
//                  -
//                  --- 3
//
// The result is returned as a single string, lines separated by "\n".

package kata

func DrawRuler(t, n int) string {
	return ""
}