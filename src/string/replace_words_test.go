package string

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_replaceWords(t *testing.T) {
	strList := []string{"aab", "bbcd", "a"}
	sort.SliceStable(strList, func(i, j int) bool {
		return len(strList[i]) > len(strList[j])
	})
	fmt.Println(strList)
}

func Test_replaceWords1(t *testing.T) {
	dictionary := []string{"cat", "bat", "rat"}
	sentence := "the cattle was rattled by the battery"
	fmt.Println(replaceWords(dictionary, sentence))
	is := assert.New(t)
	is.Equal(replaceWords(dictionary, sentence), "the cat was rat by the bat")

}
