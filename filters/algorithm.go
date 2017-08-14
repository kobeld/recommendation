package filters

import (
	"math"
)

type similarity func(ratedItems1, ratedItems2 map[string]float64) float64

func EuclideanDistance(ratedItems1, ratedItems2 map[string]float64) float64 {

	var (
		score float64
		items = mutualItems(ratedItems1, ratedItems2)
	)

	if len(items) == 0 {
		return score
	}

	for _, item := range items {
		score += math.Pow((ratedItems1[item] - ratedItems2[item]), 2)
	}

	return 1 / (1 + math.Sqrt(score))
}

func PearsonCorrelation(ratedItems1, ratedItems2 map[string]float64) float64 {
	var (
		score float64
		items = mutualItems(ratedItems1, ratedItems2)
		n     = float64(len(items))
	)

	if n == 0 {
		return score
	}

	var (
		sum1, sum2     float64
		sumSq1, sumSq2 float64
		pSum           float64
	)

	for _, item := range items {
		// Add up all the preferences
		sum1 += ratedItems1[item]
		sum2 += ratedItems2[item]

		// Sum up the squares
		sumSq1 += math.Pow(ratedItems1[item], 2)
		sumSq2 += math.Pow(ratedItems2[item], 2)

		// Sum up the products
		pSum += ratedItems1[item] * ratedItems2[item]
	}

	// Calculate Pearson score
	num := pSum - (sum1 * sum2 / n)
	den := math.Sqrt((sumSq1 - math.Pow(sum1, 2)/n) * (sumSq2 - math.Pow(sum2, 2)/n))

	if den == 0 {
		return 0
	}

	return num / den
}

// Check: https://measuringu.com/affinity-analysis/
// This algorithm basically used for caculate implicit feedback with binary data
// result = (ad - bc) / sprt((a + b)(c + d)(a + c)(b + d))
// a: both exist
// b/c: one exist
// d: both nonexist
func PhiCorrelation(ratedItems1, ratedItems2 map[string]float64) float64 {

	var (
		both, single1, single2, nboth float64
	)

	for item, value1 := range ratedItems1 {
		value2 := ratedItems2[item]

		switch {
		case value1 == 1 && value2 == 1:
			both += 1.0
		case value1 == 1 && value2 == 0:
			single1 += 1.0
		case value1 == 0 && value2 == 1:
			single2 += 1.0
		case value1 == 0 && value2 == 0:
			nboth += 1.0
		}
	}

	// fmt.Printf("Both: %+v\n", both)
	// fmt.Printf("Single1: %+v\n", single1)
	// fmt.Printf("Single2: %+v\n", single2)
	// fmt.Printf("NBoth: %+v\n", nboth)

	numerator := (both*nboth - single1*single2)
	// fmt.Printf("numerator: %+v\n", numerator)
	// fmt.Printf("denominator: %+v\n", (both+single1)*(single2+nboth)*
	// 	(both+single2)*(single1+nboth))

	denominator := math.Sqrt((both + single1) * (single2 + nboth) *
		(both + single2) * (single1 + nboth))

	return numerator / denominator
}

func mutualItems(col1, col2 map[string]float64) (items []string) {
	for key := range col1 {
		if _, ok := col2[key]; ok {
			items = append(items, key)
		}
	}
	return
}
