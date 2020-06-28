package observer

func ExampleObserver()  {

	subject := NewSubject()

	reader1 := NewReader("o1")

	reader2 := NewReader("o2")

	subject.Attach(reader1)
	subject.Attach(reader2)

	subject.UpdateContext("changed")
	// Output:
	// o1-changed1
	// o2-changed

}