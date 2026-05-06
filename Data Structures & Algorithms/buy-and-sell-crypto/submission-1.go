func maxProfit(prices []int) int {
	res, minBuy := 0, math.MaxInt32
	for _, sell := range prices{
		if sell < minBuy {
			minBuy =sell 
		}
		if sell - minBuy > res {
			res = sell- minBuy
		}
	}
	return res
}
