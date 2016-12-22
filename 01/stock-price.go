package main

import "fmt"

func main() {
	stock_prices_yesterday := []int{10, 7, 5, 8, 11, 9}

	fmt.Println(get_max_profit(stock_prices_yesterday))
	fmt.Println(get_max_profit_optimized(stock_prices_yesterday))
}

func get_max_profit(s []int) int {
	// Initial profit is zero.
	var bestProfit int
	// Loop through the slice twice.
	for firstTime, firstPrice := range s {
		// If the times are the same, break.
		// Possible optimization, if the time is less than our current time, break. Have to use c-style for loop, not for-range.
		// Rewrite comment
		for secondTime, secondPrice := range s {
			if firstTime < secondTime {
				// Compare the first loop to the second loop.
				// If first comparison (firstTime == 0, secondTime == 1), store Profit anyways, incase secondPrice - firstPrice is 0. Choosing this instead of `>=` due to readability.
				if firstTime == 0 && secondTime == 1 {
					bestProfit = secondPrice - firstPrice
					continue
				}
				if secondPrice-firstPrice > bestProfit {
					// If the profit is higher than previous profit, store profit.
					// Profit is 'the positive difference between SecondPrice and firstPrice'
					bestProfit = secondPrice - firstPrice
				}

				// Stored Profit after loop completion should be highest profit.
				// Assumption: we don't care about duplicate Profits.
			}
		}
	}

	// Possible implementations, find lowest price, and highest price after lowest price.

	return bestProfit
}

func get_max_profit_optimized(s []int) int {
	// Initial profit is zero.
	var bestProfit int
	// Loop through the slice twice.
	for firstTime := 0; firstTime < len(s); firstTime++ {
		// If the times are the same, break.
		// Possible optimization, if the time is less than our current time, break. Have to use c-style for loop, not for-range.
		// Rewrite comment
		for secondTime := firstTime + 1; secondTime < len(s); secondTime++ {
			if firstTime < secondTime {
				// Compare the first loop to the second loop.
				// If first comparison (firstTime == 0, secondTime == 1), store Profit anyways, incase secondPrice - firstPrice is 0. Choosing this instead of `>=` due to readability.
				if firstTime == 0 && secondTime == 1 {
					bestProfit = s[secondTime] - s[firstTime]
					continue
				}
				if s[secondTime]-s[firstTime] > bestProfit {
					// If the profit is higher than previous profit, store profit.
					// Profit is 'the positive difference between SecondPrice and firstTime'
					bestProfit = s[secondTime] - s[firstTime]
				}

				// Stored Profit after loop completion should be highest profit.
				// Assumption: we don't care about duplicate Profits.
			}
		}
	}

	// Possible implementations, find lowest price, and highest price after lowest price.

	return bestProfit
}
