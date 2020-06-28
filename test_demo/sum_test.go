package util

import (
	"fmt"
)

//func TestSum(t *testing.T)  {
//
//	a, b, c := 1, 2, 3
//
//	if sum(a, b) != c {
//		t.Error("error")
//	}
//}
//
//func BenchmarkSum(b *testing.B)  {
//	for i := 0; i < b.N; i++ {
//		sum(1, 2)
//	}
//
//}

func ExampleSum() {
	fmt.Println(sum(1, 4))
	// Output:
	// 5
}