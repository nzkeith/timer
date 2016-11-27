package main

import "testing"
import "io"
import "time"
import "fmt"
import "github.com/stretchr/testify/assert"
import "regexp"
import "strconv"

var regex, _ = regexp.Compile("{ ==timer== Elapsed (.+)s }")

func Test(t *testing.T) {
	reader, writer := io.Pipe()

	go func() {
		time.Sleep(400 * time.Millisecond)
		fmt.Fprintln(writer, "About 0.4 seconds in")
		time.Sleep(700 * time.Millisecond)
		fmt.Fprintln(writer, "About 1.1 seconds in")
		writer.Close()
	}()

	var lines LineReader
	run(reader, &lines)
	assert.Equal(t, 4, len(lines))
	assert.InDelta(t, 0.4, parseElapsedTime(lines[0]), 0.1)
	assert.InDelta(t, 1.1, parseElapsedTime(lines[2]), 0.1)
}

func parseElapsedTime(line string) float64 {
	match := regex.FindStringSubmatch(line)
	float, _ := strconv.ParseFloat(match[1], 64)
	return float
}

type LineReader []string

func (s *LineReader) Write(data []byte) (int, error) {
	*s = append(*s, string(data))
	return len(data), nil
}
