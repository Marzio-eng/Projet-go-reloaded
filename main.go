package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: go run . <input> <output>")
		os.Exit(1)
	}

	input, output := os.Args[1], os.Args[2]

	data, err := os.ReadFile(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, "lecture :", err)
		os.Exit(1)
	}

	text := string(data)

	lines := strings.Split(text, "\n")
	for i, line := range lines {
		lines[i] = Process(line)
	}
	result := strings.Join(lines, "\n")

	err = os.WriteFile(output, []byte(result), 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ecriture :", err)
		os.Exit(1)
	}
}
