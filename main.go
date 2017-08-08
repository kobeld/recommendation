package main

import (
	"github.com/kobeld/recommendation/algorithm"
	"github.com/kobeld/recommendation/datasource"
)

func main() {
	var (
		score       float64
		ratedItems1 = datasource.Critics["Lisa"]
		ratedItems2 = datasource.Critics["Gene"]
	)

	score = algorithm.EuclideanDistance(ratedItems1, ratedItems2)
	println(score)

	score = algorithm.PearsonCorrelation(ratedItems1, ratedItems2)
	println(score)
}
