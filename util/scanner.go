package util

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"tjz"
)

func GetString() string {
	reader := bufio.NewReader(os.Stdin)
	buff, _, err := reader.ReadLine()
	if err != nil || err == io.EOF {
		tjz.Fatal(err)
	}
	return string(buff)
}

func GetInt() int {
	reader := bufio.NewReader(os.Stdin)
	buff, _, err := reader.ReadLine()
	if err != nil || err == io.EOF {
		tjz.Fatal(err)
	}
	val, err := strconv.Atoi(string(buff))
	if err != nil {
		tjz.Fatal(err)
	}
	return val
}
