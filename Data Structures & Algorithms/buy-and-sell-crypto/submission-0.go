func maxProfit(prices []int) int {
	res, min := 0, int(1e9+7)
	for _, price := range prices{
		if price < min {
			min = price
		}
		if price - min > res {
			res = price - min
		}
	}
	return res
}
