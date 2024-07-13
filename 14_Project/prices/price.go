// 2. package prices

// import "fmt"

// type TaxIncludedPriceJob struct {
// 	TaxRate           float64
// 	InputPrices       []float64
// 	TaxIncludedPrices map[string]float64 // [{"0.1": 98}]
// }

// func (job TaxIncludedPriceJob) Process() {
// 	result := make(map[string]float64)
// 	for _, price := range job.InputPrices {
// 		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
// 	}

// 	fmt.Println(result)
// }

// func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
// 	return &TaxIncludedPriceJob{
// 		TaxRate:     taxRate,
// 		InputPrices: []float64{10, 20, 30},
// 	}
// }

package prices

import (
	"fmt"
	"tax_project/conversion"
	filemanager "tax_project/fileManager"
)

type TaxIncludedPriceJob struct {
	IOManager         filemanager.FileManager
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string // [{"0.1": 98}]
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := job.IOManager.ReadLines()

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	// file, err := os.Open("prices.txt")
	// if err != nil {
	// 	fmt.Println("Error opening file:", err)
	// 	return
	// }
	// scanner := bufio.NewScanner(file)

	// var lines []string
	// for scanner.Scan() {
	// 	lines = append(lines, scanner.Text())
	// }

	// err = scanner.Err()

	// if err != nil {
	// 	fmt.Println("Error reading file:", err)
	// 	file.Close()
	// 	return
	// }

	// prices := make([]float64, len(lines))

	// for lineIndex, line := range lines {
	// 	floatPrice, err := strconv.ParseFloat(line, 64)
	// 	if err != nil {
	// 		fmt.Printf("Error parsing line %d: %s\n", lineIndex+1, err)
	// 		file.Close()
	// 		return
	// 	}
	// 	prices[lineIndex] = floatPrice
	// }

	prices, err := conversion.StrigToFloats(lines)

	if err != nil {
		fmt.Println("Error parsing prices:", err)
		// file.Close()
		return
	}

	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrices := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrices)
	}

	job.TaxIncludedPrices = result
	job.IOManager.WriteJson(job)

	//fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64, fm filemanager.FileManager) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   fm,
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}
