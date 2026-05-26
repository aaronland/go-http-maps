package tile

import (
	"testing"
)

func TestZXYToNested(t *testing.T) {

	zoom, col, row := 10, 164, 395

	expected := "10/000/164/000/395.png"

	got := ZXYToNested(zoom, col, row)

	if got != expected {

		t.Fatalf("ZXYToNested(%d,%d,%d) = %q; expected %q", zoom, col, row, got, expected)
	}
}

func TestNestedToZXY(t *testing.T) {

	nested := "https://tileserver.example.com/10/000/164/000/395.png"

	zoom, col, row, err := NestedToZXY(nested)

	if err != nil {
		t.Fatalf("NestedToZXY(%q) returned error: %v", nested, err)
	}

	if zoom != 10 || col != 164 || row != 395 {

		t.Fatalf("NestedToZXY(%q) = (%d,%d,%d); expected (10,164,395)", nested, zoom, col, row)
	}

	got := ZXYToNested(zoom, col, row)

	if got != "10/000/164/000/395.png" {

		t.Fatalf("round‑trip mismatch: got %q; expected %q", got, "10/000/164/000/395.png")
	}
}
