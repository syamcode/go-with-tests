package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func TestRepeat(t *testing.T) {
	t.Run("Repeat 'a' character 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := strings.Repeat("a", 5)

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
	t.Run("Repeat 'c; character n times", func(t *testing.T) {
		repeated := Repeat("c", 10)
		expected := strings.Repeat("c", 10)

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}
