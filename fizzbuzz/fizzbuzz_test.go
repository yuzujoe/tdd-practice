package fizzbuzz

import "testing"

func TestConvert(t *testing.T) {
	got := Convert(1)
	if got != "1" {
		t.Errorf(`Convert(1) is %q`, got)
	}

	got = Convert(2)
	if got != "2" {
		t.Errorf(`Convert(2) is %q`, got)
	}

	got = Convert(3)
	if got != "Fizz" {
		t.Errorf(`Convert(3) is %q`, got)
	}

	got = Convert(5)
	if got != "Buzz" {
		t.Errorf(`Convert(5) is %q`, got)
	}
}
