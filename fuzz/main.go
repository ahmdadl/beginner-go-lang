package main

import "fmt"

func Reverse(s string) string {

	fmt.Printf("input: %q\n", s)
    b := []rune(s)
    fmt.Printf("runes: %q\n", b)
	// b := []byte(s)

	for start, end := 0, len(b)-1;start < len(b)/2; start,end = start+1, end-1 {
		b[start], b[end] = b[end], b[start]
	}
	
	return string(b)
}

func main() {
	input := "The quick brown fox jumped over the lazy dog"
    rev := Reverse(input)
    doubleRev := Reverse(rev)
    fmt.Printf("original: %q\n", input)
    fmt.Printf("reversed: %q\n", rev)
    fmt.Printf("reversed again: %q\n", doubleRev)
}