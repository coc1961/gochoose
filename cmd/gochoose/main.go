package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/coc1961/gochoose/internal/choose"
)

func main() {
	ch := choose.New()
	defer ch.Close()

	scanner := bufio.NewScanner(os.Stdin)

	options := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		options = append(options, line)
	}
	x, _ := ch.Choose(options)
	ch.Close()
	fmt.Print(x)
}
