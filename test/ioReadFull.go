package main

import "os"

/*
	测试连续读取的问题，对应hent/datapack_test.go
*/
import (
	"fmt"
	"io"
)

func main() {
	file, err := os.Open("test/data.txt")
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return
	}
	defer file.Close()

	buf := make([]byte, 3)
	n, err := io.ReadFull(file, buf)
	if err != nil {
		if err == io.ErrUnexpectedEOF {
			fmt.Println("Unexpected EOF encountered")
		} else {
			fmt.Println("Error reading file:", err)
		}
		return
	}

	buf2 := make([]byte, 5)
	n2, err2 := io.ReadFull(file, buf2)
	if err2 != nil {
		if err2 == io.ErrUnexpectedEOF {
			fmt.Println("Unexpected EOF encountered")
		} else {
			fmt.Println("Error reading file:", err2)
		}
		return
	}

	fmt.Printf("Read %d bytes: %s\n", n, buf)
	fmt.Printf("Read %d bytes: %s\n", n2, buf2)
}
