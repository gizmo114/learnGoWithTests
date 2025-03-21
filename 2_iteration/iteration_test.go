package iteration

import "testing"

func TestRepeat(t *testing.T) {
	get := Repeat("a", 5)
	want := "aaaaa"

	if get != want {
		t.Errorf("got %q want %q", get, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", b.N)
	}
}
