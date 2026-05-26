package tile

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// NestedToZXY converts a nested path (e.g. "10/000/164/000/395.png")
// into zoom, column, row coordinates.
func NestedToZXY(nested string) (int, int, int, error) {
	return parseNested(nested)
}

// ZXYToNested converts zoom, column, row into the nested path string.
func ZXYToNested(zoom, col, row int) string {
	return buildNested(zoom, col, row)
}

// pad left‑pads s with c until its length is n.
func pad(s string, n int, c byte) string {

	if len(s) >= n {
		return s
	}

	return strings.Repeat(string(c), n-len(s)) + s
}

// isNumeric reports whether s consists solely of digits.
func isNumeric(s string) bool {
	return regexp.MustCompile(`^\d+$`).MatchString(s)
}

// atoiJoin concatenates two numeric strings and converts the result to int.
func atoiJoin(a, b string) int {
	n, _ := strconv.Atoi(a + b)
	return n
}

// format converts an integer to a 6‑digit zero‑padded string and returns
// the nested representation “AAA/BBB”.
func format(i int) string {
	s := pad(strconv.Itoa(i), 6, '0')
	return s[:3] + "/" + s[3:]
}

// buildNested returns the “nested” URL path for the supplied zoom/row/col.
// The path is of the form “{Z}/{X1}{X2}/{Y1}{Y2}.png”.
func buildNested(zoom, col, row int) string {

	// Wrap the column value around the tile size (256×256 → 1<<zoom).

	max := 1 << zoom
	col = col % max

	if col < 0 {
		col += max
	}

	return fmt.Sprintf("%d/%s/%s.png",
		zoom,
		format(col),
		format(row),
	)
}

// parseNested extracts zoom, column and row from a nested URL path.
// The input may be a bare path or a full URL; the file extension (if any)
// is ignored.
func parseNested(raw string) (zoom, col, row int, err error) {

	path := raw

	if ext := filepath.Ext(path); ext != "" {
		path = strings.TrimSuffix(path, ext)
	}

	// Split on '/' and keep only numeric parts.

	parts := strings.FieldsFunc(path, func(r rune) bool {
		return r == '/'
	})

	num := make([]string, 0, len(parts))

	for _, p := range parts {

		if isNumeric(p) {
			num = append(num, p)
		}
	}

	if len(num) < 5 {
		return 0, 0, 0, fmt.Errorf("not enough numeric segments to infer Z/X/Y")
	}

	// In the nested format we expect 5 numeric parts:
	//   Z / X1X2 / Y1Y2
	// (X1X2 and Y1Y2 are 6‑digit strings split into two 3‑digit chunks).

	zoom, _ = strconv.Atoi(num[len(num)-5])
	col = atoiJoin(num[len(num)-4], num[len(num)-3])
	row = atoiJoin(num[len(num)-2], num[len(num)-1])

	return zoom, col, row, nil
}
