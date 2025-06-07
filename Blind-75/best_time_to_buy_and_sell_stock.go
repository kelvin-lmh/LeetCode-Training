package blind_75

func maxProfit(prices []int) int {
	buyPrice := prices[0]
	maxProfit := 0

	for _, price := range prices[1:] {
		if maxProfit < price-buyPrice {
			maxProfit = price - buyPrice
		}
		if price < buyPrice {
			buyPrice = price
		}
	}
	return maxProfit
}
