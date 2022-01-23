package leetcode

import "testing"

func TestStockPrice_Update(t *testing.T) {
	stockPrice := Constructor()
	stockPrice.Update(45, 9233)
	stockPrice.Update(45, 9651)
	stockPrice.Update(37, 3902)
	stockPrice.Update(25, 4833)
	stockPrice.Update(53, 4521)
	stockPrice.Update(22, 1161)
	stockPrice.Update(55, 6897)
	stockPrice.Update(20, 5354)
	stockPrice.Update(30, 5623)
	stockPrice.Update(25, 2725)
}
