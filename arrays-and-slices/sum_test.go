package arrays_and_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum collection of any size", func(t *testing.T) {
		numbers := []int{5, 6, 7, 8}

		got := Sum(numbers)
		want := 26

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	assertSums := func(t testing.TB, got, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("correctly sum tails of non empty collections", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{3, 5, 9})
		want := []int{2, 9}
		assertSums(t, got, want)
	})

	t.Run("correctly sum tails of empty collections", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{}, []int{1, 5, 15})
		want := []int{2, 15}
		assertSums(t, got, want)
	})
}
