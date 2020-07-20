package leetcode

import (
	"bytes"
	"strconv"
)

func decodeString(s string) string {
	var numStack [][]byte
	var letterStack [][]byte
	var result bytes.Buffer
	stacklen := -1

	for i := 0; i < len(s); i++ {
		if isNumber(s[i]) {
			// number to stack
			if len(numStack)-1 <= stacklen {
				numStack = append(numStack, []byte{s[i]})
				letterStack = append(letterStack, []byte{})
			} else {
				numStack[stacklen+1] = append(numStack[stacklen+1], s[i])
			}
		} else if s[i] == '[' {
			stacklen++
		} else if s[i] == ']' {
			n := len(numStack) - 1
			topNum, _ := strconv.Atoi(string(numStack[n]))
			for i := 0; i < topNum; i++ {
				if stacklen == 0 {
					result.Write(letterStack[n])
				} else {
					letterStack[n-1] = append(letterStack[n-1], letterStack[n]...)
				}
			}

			// pop
			numStack = numStack[:n]
			letterStack = letterStack[:n]
			stacklen--
		} else {
			if stacklen == -1 {
				result.WriteByte(s[i])
			} else {
				letterStack[stacklen] = append(letterStack[stacklen], s[i])
			}
		}
	}

	return result.String()
}

func isNumber(c byte) bool {
	return c >= '0' && c <= '9'
}
