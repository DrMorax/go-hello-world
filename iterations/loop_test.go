package iterations

import "testing"

func TestLoop(t *testing.T) {
	repeated := Loop("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("got %q expected %q", repeated, expected)
	}
}

func ExampleLoop() {
	repeated := "a"
	repeatCount := 5

	fm.Println(Loop(repeated, repeatCount))
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Loop("a", 5)
	}
}