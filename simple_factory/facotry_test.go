package simple_factory

import "fmt"

func ExampleFactory() {

	fmt.Println(NewTranslator("en").translate())
	fmt.Println(NewTranslator("cn").translate())

	// Output:
	// en
	// cn
}
