package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	n := Years(10)
	if n != 70 {
		t.Error("Got", n, "want", 70)
	}
}

func TestYearsTwo(t *testing.T) {
	n := YearsTwo(10)
	if n != 70 {
		t.Error("Got", n, "want", 70)
	}
}

func exampleYears() {
	fmt.Println(Years(10))
	// Output:
	// 70
}

func exampleYearsTwo() {
	fmt.Println(YearsTwo(10))
	// Output:
	// 70
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}
