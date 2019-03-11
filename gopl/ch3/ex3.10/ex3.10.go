package main

import (
	"bytes"
	"fmt"
	"os"
)

func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	if n <= 3 {
		return s
	}

	if i := n % 3; i == 0 {
		for j := 0; j < n/3; j++ {
			if j > 0 {
				buf.WriteString(",")
			}
			fmt.Fprintf(&buf, "%v", s[j*3:(j+1)*3])
		}
	} else {
		for j := 0; j < n/3; j++ {
			if j == 0 {
				fmt.Fprintf(&buf, "%v,", s[:i])
			} else {
				buf.WriteString(",")
			}
			fmt.Fprintf(&buf, "%v", s[j*3+i:(j+1)*3+i])
		}
	}
	return buf.String()
}

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}
