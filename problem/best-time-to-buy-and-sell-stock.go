package problem

func MaxProfit(prices []int) int {
	buy := prices[0]
	profit := 0
	for i := range prices {
		if buy > prices[i] {
			buy = prices[i]
		} else {
			takeProfit := prices[i] - buy
			if takeProfit > profit {
				profit = takeProfit
			}
		}
	}

	return profit
}
