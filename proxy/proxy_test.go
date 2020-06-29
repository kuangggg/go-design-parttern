package proxy

import (
	"fmt"
)

//func TestProxy(t *testing.T) {
//
//	s := &Proxy{}
//
//	res := s.Do()
//
//	println(res)
//
//	if res != "pre:real:after" {
//		t.Fail()
//	}
//
//}

func ExampleProxy_Do() {

	s := &Proxy{}

	res := s.Do()

	fmt.Println(res)
	// Output:
	// before:real:after

}
