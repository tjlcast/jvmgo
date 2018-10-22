package main

import (
	"bytes"
	"fmt"
)

/**
focus:
buf
fmt -> Fprintf, println, Sprintf
strconv
 */

func ints2String(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')

	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	fmt.Println([]int{1, 2, 3, 4, 5})
}
