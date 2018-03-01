// UVa 444 - Encoder and Decoder

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func encode(line string) string {
	var midBuilder strings.Builder
	for i := range line {
		midBuilder.WriteString(strconv.Itoa(int(line[i])))
	}
	var ret string
	mid := midBuilder.String()
	for i := range mid {
		ret = string(mid[i]) + ret
	}
	return ret
}
func decode(line string) string {
	var ret string
	ascii, base := 0, 1
	for i := range line {
		ascii += int(line[i]-'0') * base
		base *= 10
		if ascii >= 32 {
			ret = string(ascii) + ret
			ascii, base = 0, 1
		}
	}
	return ret
}

func main() {
	in, _ := os.Open("444.in")
	defer in.Close()
	out, _ := os.Create("444.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		if line := s.Text(); line[0] >= '0' && line[0] <= '9' {
			fmt.Fprintln(out, decode(line))
		} else {
			fmt.Fprintln(out, encode(line))
		}
	}
}
