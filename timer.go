package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	run(os.Stdin, os.Stdout)
}

func run(input io.Reader, output io.Writer) {
	scanner := bufio.NewScanner(input)
	start := time.Now()
	for scanner.Scan() {
		elapsed := time.Since(start)
		fmt.Fprintf(output, "{ ==timer== Elapsed %.3fs }\n", elapsed.Seconds())
		fmt.Fprintln(output, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
