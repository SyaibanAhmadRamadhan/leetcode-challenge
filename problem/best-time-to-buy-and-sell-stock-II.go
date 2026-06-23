package problem

func MaxProfitII(prices []int) int {
	buy := prices[0]
	profit := 0

	for i := 0; i < len(prices); i++ {
		if buy > prices[i] {
			buy = prices[i]
		} else {
			takeProfit := prices[i] - buy
			if takeProfit > 0 {
				profit += takeProfit
			}
			buy = prices[i]
		}
	}

	// solusi lain, cara 1 sebenernya buy itu representasi dari harga dihari ini
	for i := 0; i < len(prices); i++ {
		if i+1 < len(prices) {
			takeProfit := prices[i+1] - prices[i]
			if takeProfit > 0 {
				profit += takeProfit
			}
		}
	}

	return profit
}
