package leetcode

import (
	"bytes"
	"sort"
	"strconv"
)

type cnt = map[string]int

func countOfAtoms(formula string) string {
	var result bytes.Buffer
	formulaStack := []cnt{}
	formulaStack = append(formulaStack, make(cnt))
	stacklen := 0

	for i := 0; i < len(formula); {
		c := formula[i]
		if isUpper(c) || c == ')' {
			i++
			atom := []byte{c}
			times := []byte{}
			for i < len(formula) {
				if isLower(formula[i]) {
					atom = append(atom, formula[i])
				} else if isNum(formula[i]) {
					times = append(times, formula[i])
				} else {
					break
				}
				i++
			}
			t := 1
			if len(times) > 0 {
				t, _ = strconv.Atoi(string(times))
			}

			if isUpper(c) {
				formulaStack[stacklen][string(atom)] += t
			} else {
				for k, v := range formulaStack[stacklen] {
					formulaStack[stacklen][k] = v * t
				}

				for k, v := range formulaStack[stacklen] {
					formulaStack[stacklen-1][k] += v
				}
				if stacklen >= 1 {
					formulaStack = formulaStack[:stacklen]
				}
				stacklen--
			}
		} else if c == '(' {
			stacklen++
			formulaStack = append(formulaStack, make(cnt))
			i++
		}
	}

	keys := make([]string, 0, len(formulaStack[0]))
	for k, _ := range formulaStack[0] {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		if v, ok := formulaStack[0][k]; ok {
			result.WriteString(k)
			if v != 1 {
				result.WriteString(strconv.Itoa(v))
			}
		}
	}
	return result.String()
}

func isNum(c byte) bool {
	return c >= '0' && c <= '9'
}

func isUpper(c byte) bool {
	return c >= 'A' && c <= 'Z'
}

func isLower(c byte) bool {
	return c >= 'a' && c <= 'z'
}
