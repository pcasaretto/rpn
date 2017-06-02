package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/pcasaretto/rpn"
)

func main() {
	stack := &rpn.Stack{}
	scanner := bufio.NewScanner(os.Stdout)
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "+" {
			rpn.Add(stack)
		} else {
			f, err := strconv.ParseFloat(text, 64)
			if err != nil {
				panic(err)
			}
			stack.Push(f)
		}
		fmt.Println(stack)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}
