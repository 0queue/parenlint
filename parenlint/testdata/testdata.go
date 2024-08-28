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

	fmt.Println("a", // want `Argument on same line as left paren`
		"b", "c", "d", // want `Argument on same line as previous argument` `Argument on same line as previous argument`
		"e") // want `Argument on same line as right paren`

	fmt.Println(
		"a" + // want `Argument on same line as right paren`
			"b")

	fmt.Println(
		"a"+
			"b", "c", // want `Argument on same line as previous argument`
	)

	fmt.Println("hanging func?", func() string {
		return "all good"
	}())
}
