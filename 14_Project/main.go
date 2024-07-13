//1. package main

// import "fmt"

// func main() {
// 	prices := []float64{10, 20, 30}
// 	taxRates := []float64{0, 0.07, 0.1, 0.15}

// 	result := make(map[float64][]float64)
// 	// [{0.07 : {10, 20, 30}}]
// 	for _, taxRate := range taxRates {
// 		taxIncludePrices := make([]float64, len(prices))
// 		for priceIndex, price := range prices {
// 			taxIncludePrices[priceIndex] = price * (1 + taxRate)
// 		}
// 		result[taxRate] = taxIncludePrices
// 	}

// 	fmt.Println(result)
// }

package main

import (
	"fmt"
	filemanager "tax_project/fileManager"
	"tax_project/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	result := make(map[float64][]float64)
	// [{0.07 : {10, 20, 30}}]
	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(taxRate, fm)
		priceJob.Process()
	}

	fmt.Println(result)
}
