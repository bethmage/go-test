// Sample _test appends to the pacakage name
// Sample function to displayed in godoc GUI
// BUT... this will not run in godoc GUI playground, because package: github.com/bethmage/go-test/morestrings is NOT public availabe!
package morestrings_test

import (
	"fmt"

	"github.com/bethmage/go-test/morestrings"
)

// Test function preFixed with Example.
//
// Sample Output: comment.
//  As it executes the example, the testing framework captures data written to standard output
//  and then compares the output against the example's "Output:" comment.
//  The test passes if the test's output matches its output comment.
func ExampleReverseRunes() {
	fmt.Println(morestrings.ReverseRunes("hello"))
	// Output: olleh
}
