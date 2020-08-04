package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/coc1961/gochoose/internal/choose"
)

func main() {
	ch := choose.New()

	options := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		options = append(options, line)
	}
	os.Stdin.Close()
	x, err := ch.Choose(options)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Print(x)
}
