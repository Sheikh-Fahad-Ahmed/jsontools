package helper

import (
	"bufio"
	"fmt"
)

func prompt(reader *bufio.Reader, label string) string {
	fmt.Print(label)
	line, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}
	return string(line)
}
