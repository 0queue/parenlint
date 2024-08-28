package examples

import "fmt"

func stuff() {
	fmt.Println("this is fine")

	fmt.Println(
		"this",
		"is",
		"also",
		"fine",
	)

	fmt.Println("a", // want `Single line function call with arguments on multiple lines`
		"b", "c", "d",
		"e")

	fmt.Println( // want `Multiline function call with multiple arguments on single line`
		"a" +
			"b")

	fmt.Println( // want `Multiline function call with multiple arguments on single line`
		"a"+
			"b", "c",
	)

	fmt.Println("hanging func?", func() string {
		return "all good"
	}())

	fmt.Println("weird but okay", func() string {
		return "string"
	}, "I guess", func() string {
		return "string"
	}())
}
