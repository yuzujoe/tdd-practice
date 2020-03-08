package fizzbuzz

import "testing"

func TestConvert(t *testing.T) {
	got := Convert(1)
	if got != "1" {
		t.Errorf(`Convert(1) is %q`, got)
	}
}
