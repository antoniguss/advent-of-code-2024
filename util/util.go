package util

import (
	"bufio"
	"os"
)

func Freeze() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
}
