package main

// see go.mod file
import (
	"fmt"
        "github.com/bethmage/go-test/morestrings"
)

func main() {
        var str string  = "!oG ,olleH"
        fmt.Printf("-- DO reverse and print: %s\n", str) 
	fmt.Println(morestrings.ReverseRunes(str))
}

