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

func TopMatches(prefs datasource.Prefs, person string, algorithm similarity) (ps []PersonScore) {

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

		personScore := PersonScore{
			Name:  name,
			Score: algorithm(personItems, otherItems),
		}

		ps = append(ps, personScore)
	}

	sort.Sort(ByScore(ps))

	return
}
