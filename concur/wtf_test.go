package concur

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestUt(t *testing.T) {
	s1 := "aaa"
	s2 := "💺👍✓"
	extr(s1)
	extr(s2)

}

func extr(s1 string) {
	fmt.Printf("str: %v len:%v", s1, utf8.RuneCountInString(s1))
}
