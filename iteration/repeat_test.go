package iteration_test

import (
	"testing"

	"github.com/sophiabrandt/learn-go-with-tests/iteration"
)

func TestRepeat(t *testing.T) {
	repeated := iteration.Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iteration.Repeat("a", 8)
	}
}
