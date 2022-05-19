package back_track

import (
	"fmt"
	"testing"
)

func Test_countVowels(t *testing.T) {
	sum := countVowels("abc")
	fmt.Println(sum)
}

func Test_countVowels2(t *testing.T) {
	countVowels("aba")
}
