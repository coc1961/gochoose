package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/coc1961/gochoose/internal/choose"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	options := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		options = append(options, line)
	}
	if len(options) == 0 {
		return
	}

	ch := choose.New()
	if len(os.Args) > 1 {
		ch.SetSelected(os.Args[1])
	}
	x, err := ch.Choose(options)
	if err != nil {
		os.Exit(1)
	}
	fmt.Print(x)
}
