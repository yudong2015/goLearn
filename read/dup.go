package read

import (
	"bufio"
	"fmt"
	"os"
)

// ReadDup : prints the text of each line that appears more
//than once in the standard input, preceded by its count.
func ReadDup() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "" {
			break
		}
		counts[input.Text()]++
	}

	for text, n := range counts {
		fmt.Printf("%d\t%s\n", n, text)
	}
}
