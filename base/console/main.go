package console

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var reader = bufio.NewReader(os.Stdin)

func ReadString(a ...interface{}) string {
	if len(a) > 0 {
		Write(a)
	}
	result, _, _ := reader.ReadLine()
	return string(result)
}

func ReadInt(a ...interface{}) int {
	result, _ := strconv.Atoi(ReadString(a...))
	return result
}

func ReadFloat(a ...interface{}) float64 {
	result, _ := strconv.ParseFloat(ReadString(a...), 10)
	return result
}

func Write(a ...interface{}) {
	fmt.Print(a...)
}

func Writeln(a ...interface{}) {
	fmt.Println(a...)
}
