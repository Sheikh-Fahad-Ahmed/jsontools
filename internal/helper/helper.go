package helper

import (
	"bufio"
	"fmt"
)

func Prompt(reader *bufio.Reader, label string) string {
	fmt.Print(label)
	line, _, err := reader.ReadLine()
	CheckErr(err)
	
	return string(line)
}

func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}
