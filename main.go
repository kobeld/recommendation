package main

import (
	"fmt"

	"github.com/kobeld/recommendation/datasource"
	"github.com/kobeld/recommendation/filters"
)

func main() {
	var (
		score   float64
		person1 = "Lisa"
		person2 = "Michael"
	)

	fmt.Printf("1. Getting score for %s and %s:\n", person1, person2)

	score = filters.GetScore(datasource.BoughtItems, person1, person2, filters.EuclideanDistance)
	fmt.Printf("Euclidean distance score is: %+v \n", score)

	score = filters.GetScore(datasource.BoughtItems, person1, person2, filters.PearsonCorrelation)
	fmt.Printf("Pearson Correlation score is: %+v \n", score)

	var (
		targetPerson = "Toby"
		targetMovie  = "Superman Returns"
	)

	fmt.Printf("\n2. Ranking the Critics for %s\n", targetPerson)
	personSocres := filters.TopMatches(datasource.BoughtItems, targetPerson, filters.PearsonCorrelation)
	for i, ps := range personSocres {
		fmt.Printf("  %d) %s: %+v\n", i+1, ps.Name, ps.Score)
	}

	fmt.Printf("\n3. Getting recommendations for %s\n", targetPerson)
	itemScores := filters.GetRecommendations(datasource.BoughtItems, targetPerson, filters.EuclideanDistance)
	for i, ps := range itemScores {
		fmt.Printf("  %d) %s: %+v\n", i+1, ps.Name, ps.Score)
	}

	fmt.Printf("\n4. Matching products for \"%s\"\n", targetMovie)
	moviePrefs := filters.TransformPrefs(datasource.BoughtItems)
	productScores := filters.TopMatches(moviePrefs, targetMovie, filters.PearsonCorrelation)
	for i, ps := range productScores {
		fmt.Printf("  %d) %s: %+v\n", i+1, ps.Name, ps.Score)
	}

	// targetMovie = "Just My Luck"
	fmt.Printf("\n5. Getting recommended critics for \"%s\"\n", targetMovie)
	itemScores = filters.GetRecommendations(moviePrefs, targetMovie, filters.PearsonCorrelation)
	for i, ps := range itemScores {
		fmt.Printf("  %d) %s: %+v\n", i+1, ps.Name, ps.Score)
	}

	fmt.Printf("\n6. Item-based filtering, Getting recommendations for %s\n", targetPerson)
	itemMatch := filters.CalculateSimilarItems(moviePrefs)
	userRates := datasource.BoughtItems[targetPerson]
	itemScores = filters.GetRecommendationItems(itemMatch, userRates)
	for i, ps := range itemScores {
		fmt.Printf("  %d) %s: %+v\n", i+1, ps.Name, ps.Score)
	}

	bookPrefs := filters.TransformPrefs(datasource.BoughtBooks)
	scoreBookAB := filters.PhiCorrelation(bookPrefs["Book A"], bookPrefs["Book B"])
	fmt.Printf("\n7. Phi Correlation for ‘%s’ and '%s' is %+v\n", "Book A", "Book B", scoreBookAB)

}
