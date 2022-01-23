package leetcode

import "testing"

func TestMyAtoi(t *testing.T) {
	testCase := []string{
		"    -32",
		"sgdsg dsagdsg 432",
		"00034",
		"34546576768",
		"  -+31",
		"45x96",
		"00000-42a1234",
		"  0000000000012345678",
	}

	for _, v := range testCase {
		t.Log(myAtoi(v))
	}

}
