package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	got := FizzBuzz(1)
	if got != "1" {
		t.Errorf(`FizzBuzz(1) is %q`, got)
	}
}
