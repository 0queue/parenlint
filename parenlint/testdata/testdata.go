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

func thisisfine(a int) {
	fmt.Println(a)
}

func thisisalsofine(
	a int,
	b string,
) {
	fmt.Println(a, b)
}

func thisisnotfine(b int, // want `Single line function type with arguments on multiple lines`
) {
	fmt.Println(b)
}

func multisingleline( // want `Multiline function call with multiple arguments on single line`
	a int, b string,
) {
	fmt.Println(a, b)
}

func longtype(a int, b struct {
	i int
}, c string) {
	fmt.Println(a, b, c)
}

func youthoughtIforgotaboutfuncexprs() {
	a := func(a int) {}

	_ = a

	b := func(
		a int,
		b int,
	) {
	}

	_ = b

	c := func(a int, // want `Single line function call with arguments on multiple lines`
		b int) {
	}

	_ = c

	d := func( // want `Multiline function call with multiple arguments on single line`
		a int, b int,
	) {
	}

	_ = d

	e := func(string, struct {
		i int
	}, int) {
	}

	_ = e
}
