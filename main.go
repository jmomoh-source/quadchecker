package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func readAll() string {
	r := bufio.NewReader(os.Stdin)
	s, _ := r.ReadString(0)
	return s
}

func dims(s string) (int, int) {
	s = strings.TrimRight(s, "\n")
	lines := strings.Split(s, "\n")
	return len(lines[0]), len(lines)
}

func quadA(w, h int) string { return quad(w, h, 'o', 'o', 'o', 'o', '-', '|') }
func quadB(w, h int) string { return quad(w, h, '/', '\\', '\\', '/', '*', '*') }
func quadC(w, h int) string { return quad(w, h, 'A', 'A', 'C', 'C', 'B', 'B') }
func quadD(w, h int) string { return quad(w, h, 'A', 'C', 'A', 'C', 'B', 'B') }
func quadE(w, h int) string { return quad(w, h, 'A', 'C', 'C', 'A', 'B', 'B') }

func quad(w, h int, tl, tr, bl, br, hE, vE rune) string {
	if w <= 0 || h <= 0 {
		return ""
	}

	var b strings.Builder

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			switch {
			case y == 0 && x == 0:
				b.WriteRune(tl)
			case y == 0 && x == w-1:
				b.WriteRune(tr)
			case y == h-1 && x == 0:
				b.WriteRune(bl)
			case y == h-1 && x == w-1:
				b.WriteRune(br)
			case y == 0 || y == h-1:
				b.WriteRune(hE)
			case x == 0 || x == w-1:
				b.WriteRune(vE)
			default:
				b.WriteRune(' ')
			}
		}
		b.WriteRune('\n')
	}

	return b.String()
}

func main() {
	in := readAll()
	w, h := dims(in)

	matches := []string{}

	if in == quadA(w, h) { matches = append(matches, "quadA") }
	if in == quadB(w, h) { matches = append(matches, "quadB") }
	if in == quadC(w, h) { matches = append(matches, "quadC") }
	if in == quadD(w, h) { matches = append(matches, "quadD") }
	if in == quadE(w, h) { matches = append(matches, "quadE") }

	if len(matches) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	sort.Strings(matches)

	for i := range matches {
		matches[i] = "[" + matches[i] + "]"
	}

	fmt.Printf("%s [%d] [%d]\n", strings.Join(matches, " || "), w, h)
}
