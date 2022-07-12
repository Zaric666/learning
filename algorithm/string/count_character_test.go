package string

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestCount(t *testing.T) {
	reader := bufio.NewReader(os.Stdin)
	str1, _, _ := reader.ReadLine()
	str2, _, _ := reader.ReadLine()
	cnt := 0
	for _, c := range strings.ToLower(string(str1)) {
		fmt.Println(c, str2)
		if c < 1 || c > 1000 {
			continue
		}
		str2C := []rune(string(str2))

		if c == str2C[0] {
			cnt++
		}
	}
	fmt.Println(cnt)
}
