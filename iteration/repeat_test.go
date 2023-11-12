package iteration

import (
	"fmt"
	"testing"
)

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}

func ExampleRepeat() {
	repeated := Repeat("xyz", 5)
	fmt.Println(repeated)
	// Output: xyzxyzxyzxyzxyz
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("x", 3)
	expected := "xxx"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
