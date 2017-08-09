package filters

import (
	"sort"

	"github.com/kobeld/recommendation/datasource"
)

func GetScore(prefs datasource.Prefs, person1, person2 string, algorithm similarity) float64 {
	var (
		ratedItems1 = prefs[person1]
		ratedItems2 = prefs[person2]
	)

	return algorithm(ratedItems1, ratedItems2)
}

func TopMatches(prefs datasource.Prefs, person string, algorithm similarity) (ps []ItemScore) {

	var (
		personItems = prefs[person]
	)

	if len(personItems) == 0 {
		return
	}

	for name, otherItems := range prefs {
		if name == person {
			continue
		}

		itemScore := ItemScore{
			Name:  name,
			Score: algorithm(personItems, otherItems),
		}

		ps = append(ps, itemScore)
	}

	sort.Sort(ByScore(ps))

	return
}

func GetRecommendations(prefs datasource.Prefs, person string, algorithm similarity) (ps []ItemScore) {

	var (
		personItems = prefs[person]
		sumSim      = map[string]float64{}
		sumScore    = map[string]float64{}
	)

	for other, otherItems := range prefs {
		// Don't Compare me to myself
		if other == person {
			continue
		}

		// Ignore scores of zero or lower
		weight := algorithm(personItems, otherItems)
		if weight <= 0 {
			continue
		}

		for name, score := range otherItems {
			// Only score items I haven't touch yet
			if _, ok := personItems[name]; ok {
				continue
			}

			sumSim[name] += +weight
			sumScore[name] += score * weight
		}
	}

	for name, score := range sumScore {
		itemScore := ItemScore{
			Name:  name,
			Score: score / sumSim[name],
		}

		ps = append(ps, itemScore)
	}

	sort.Sort(ByScore(ps))

	return
}

func GetRecommendationItems(itemMatch datasource.Prefs, userRates map[string]float64) (ps []ItemScore) {

	var (
		sumSim   = map[string]float64{}
		sumScore = map[string]float64{}
	)

	// Loop over items rated by the user
	for ratedName, ratedScore := range userRates {

		// Loop over items similar to this one
		for name, score := range itemMatch[ratedName] {
			// Ignore if this user has already rated this item
			if _, ok := userRates[name]; ok {
				continue
			}

			sumSim[name] += score
			sumScore[name] += score * ratedScore

		}
	}

	for name, score := range sumSim {
		itemScore := ItemScore{
			Name:  name,
			Score: sumScore[name] / score,
		}

		ps = append(ps, itemScore)
	}

	sort.Sort(ByScore(ps))

	return
}
