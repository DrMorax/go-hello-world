package iterations

import (
	"testing"
	"fmt"
)

func TestLoop(t *testing.T) {
	got := Loop("a", 5)
	expected := "aaaaa"

	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}
}

func ExampleLoop() {
	got := "a"
	repeatCount := 5

	fmt.Println(Loop(got, repeatCount))
	// Output: aaaaa
}

func BenchmarkLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Loop("a", 5)
	}
}