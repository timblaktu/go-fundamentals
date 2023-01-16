package sstrings

import "fmt"

// Capitalize function to automatically export it
func Examples() {
	// 	fmt.Println(`
	// integers overflow at compile-time if assigned initial value exceeds capacity of type
	// The below produces this error:
	//   ./intest.go:7:23: cannot use 256 (untyped int constant) as uint8 value in variable declaration (overflows)
	//   ./intest.go:10:6: maxUint8 redeclared in this block
	//      ./intest.go:7:6: other declaration of maxUint8
	// var maxUint8 uint8 = 256
	// fmt.Println("value:", maxUint8)`)

	fmt.Println(`
integers wraparound at runtime if no overflow occurs at compile time`)
	var maxUint8 uint8 = 255
	fmt.Println("value:", maxUint8)
	maxUint8 += 10
	fmt.Println("value:", maxUint8)

	fmt.Println(`
Double-quoted strings interpret control chars and escape sequences /n/t.
Can contain anything except hard-newline or unescaped double-quote.`)
	a := "Say \"hello\"\n\t\tto Go!\n\n\nHi!"
	fmt.Println(a)

	fmt.Println(`
backtick-quoted strings are raw literal strings, like heredocs.
Backtick strings are good for multiline strings.`)
	b := `Say "hello" to Go!`
	fmt.Println(b)
	c := `# json data for testing
	{
		"id": 1,
		"name": "Janis",
		"email": "pearl@example.com"
	}
	`
	fmt.Println(c)

	fmt.Println(`
A single-quoted 'rune' is an alias for int_32 and is used to represent a single character.
can be made up of 1 to 3 int32 values. This allows for both single and multibyte characters.
When printed as a value, a rune will be printed as an int32 corresponding to the UTF-8 encoding of the character.
Here, we initialize a rune to 'R' and print it with %v and %T:`)
	r := 'R'
	fmt.Printf("%v (%T)\n", r, r)

	fmt.Println(`
Iterating over a utf-8 string’s bytes, like a byte array, will produce bad results
bc each utf-8 char may be multiple bytes.
Below we do this for the string: Hello, 世界`)
	utf8_hello := "Hello, 世界" // 9 characters (including the space and comma)
	for i := 0; i < len(utf8_hello); i++ {
		fmt.Printf("%d: %s\n", i, string(utf8_hello[i]))
	}
	fmt.Println(`
Proper string iteration uses the range keyword in the loop. The range keyword ensures that we
use the proper index and length of int32s to capture the proper rune value.
Below we do this for the string: Hello, 世界`)
	for i, c := range utf8_hello {
		fmt.Printf("%d: %s\n", i, string(c))
	}

	fmt.Println(`
Literal and Interpreted strings may contain utf-8 encoded non-ascii, unicode characters.`)
	utf8_string := "Goodbye, 世界"
	fmt.Println(utf8_string)
}
