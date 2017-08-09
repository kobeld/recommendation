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
		person2 = "Gene"
	)

	fmt.Printf("1. Getting score for %s and %s:\n", person1, person2)

	score = filters.GetScore(datasource.Critics, person1, person2, filters.EuclideanDistance)
	fmt.Printf("Euclidean distance score is: %+v \n", score)

	score = filters.GetScore(datasource.Critics, person1, person2, filters.PearsonCorrelation)
	fmt.Printf("Pearson Correlation score is: %+v \n", score)

	var (
		targetPerson = "Toby"
	)

	fmt.Printf("\n2. Ranking the Critics for %s\n", targetPerson)
	personSocres := filters.TopMatches(datasource.Critics, targetPerson, filters.PearsonCorrelation)
	for i, ps := range personSocres {
		fmt.Printf("  %d) %s: %+v\n", i+1, ps.Name, ps.Score)
	}

	fmt.Printf("\n3. Getting recommendations for %s\n", targetPerson)
	itemScores := filters.GetRecommendations(datasource.Critics, targetPerson, filters.PearsonCorrelation)
	for i, ps := range itemScores {
		fmt.Printf("  %d) %s: %+v\n", i+1, ps.Name, ps.Score)
	}

}
