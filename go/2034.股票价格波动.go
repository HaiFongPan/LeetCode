package leetcode

/*
 * @lc app=leetcode.cn id=2034 lang=golang
 *
 * [2034] 股票价格波动
 */

// @lc code=start
// type StockPrice struct {
// 	prices       *redblacktree.Tree
// 	timePriceMap map[int]int
// 	maxTimestamp int
// }
// func Constructor() StockPrice {
// 	return StockPrice{redblacktree.NewWithIntComparator(), map[int]int{}, 0}
// }

// func (sp *StockPrice) Update(timestamp, price int) {
// 	if prevPrice := sp.timePriceMap[timestamp]; prevPrice > 0 {
// 		if times, _ := sp.prices.Get(prevPrice); times.(int) > 1 {
// 			sp.prices.Put(prevPrice, times.(int)-1)
// 		} else {
// 			sp.prices.Remove(prevPrice)
// 		}
// 	}
// 	times := 0
// 	if val, ok := sp.prices.Get(price); ok {
// 		times = val.(int)
// 	}
// 	sp.prices.Put(price, times+1)
// 	sp.timePriceMap[timestamp] = price
// 	if timestamp >= sp.maxTimestamp {
// 		sp.maxTimestamp = timestamp
// 	}
// }

// func (sp *StockPrice) Current() int { return sp.timePriceMap[sp.maxTimestamp] }
// func (sp *StockPrice) Maximum() int { return sp.prices.Right().Key.(int) }
// func (sp *StockPrice) Minimum() int { return sp.prices.Left().Key.(int) }

/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */
// @lc code=end
