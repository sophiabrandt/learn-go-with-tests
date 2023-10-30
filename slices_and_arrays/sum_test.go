package slicesandarrays_test

import (
	"reflect"
	"testing"

	slicesandarrays "github.com/sophiabrandt/learn-go-with-tests/slices_and_arrays"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {

		numbers := []int{1, 2, 3, 4}

		got := slicesandarrays.Sum(numbers)
		want := 10

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {

	got := slicesandarrays.SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checksums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := slicesandarrays.SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checksums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := slicesandarrays.SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checksums(t, got, want)
	})
}
