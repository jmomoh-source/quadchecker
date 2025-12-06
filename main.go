package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := read()
	if input == "" {
		fmt.Println("Not a quad function")
		return
	}

	lines := strings.Split(strings.TrimRight(input, "\n"), "\n")
	h := len(lines)
	w := len(lines[0])

	quads := map[string]string{
		"quadA": quad(w, h, 'o', 'o', 'o', 'o', '-', '|'),
		"quadB": quad(w, h, '/', '\\', '\\', '/', '*', '*'),
		"quadC": quad(w, h, 'A', 'A', 'C', 'C', 'B', 'B'),
		"quadD": quad(w, h, 'A', 'C', 'A', 'C', 'B', 'B'),
		"quadE": quad(w, h, 'A', 'C', 'C', 'A', 'B', 'B'),
	}

	var matches []string
	for name, shape := range quads {
		if shape == input {
			matches = append(matches, fmt.Sprintf("[%s] [%d] [%d]", name, w, h))
		}
	}

	if len(matches) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	fmt.Println(strings.Join(matches, " || "))
}

// Read all STDIN
func read() string {
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString(0)
	if len(data) == 0 {
		b, _ := os.ReadFile("/dev/stdin")
		return string(b)
	}
	return data
}

// Quad generator
func quad(w, h int, tl, tr, bl, br, hE, vE rune) string {
	if w <= 0 || h <= 0 {
		return ""
	}

	var out strings.Builder

	for r := 1; r <= h; r++ {
		for c := 1; c <= w; c++ {
			switch {
			case r == 1 && c == 1:
				out.WriteRune(tl)
			case r == 1 && c == w:
				out.WriteRune(tr)
			case r == h && c == 1:
				out.WriteRune(bl)
			case r == h && c == w:
				out.WriteRune(br)
			case r == 1 || r == h:
				out.WriteRune(hE)
			case c == 1 || c == w:
				out.WriteRune(vE)
			default:
				out.WriteRune(' ')
			}
		}
		out.WriteRune('\n')
	}
	return out.String()
}
