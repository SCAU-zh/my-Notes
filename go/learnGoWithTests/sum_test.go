package learnGoWithTests

import (
	"reflect"
	"testing"
)
func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		number := []int{1, 2, 3, 4, 5}
		got := Sum(number)
		want := 15

		if want != got {
			t.Errorf("got %d but want %d", got, want)
		}
	})

	t.Run("collection of any numbers", func(t *testing.T) {
		number := []int{1, 2, 3, 4, 5}
		got := Sum(number)
		want := 15

		if want != got {
			t.Errorf("got %d but want %d", got, want)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{3, 4})
	want := []int{3, 7}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d but want %d", got, want)
	}
}