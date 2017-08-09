package filters

import (
	"github.com/kobeld/recommendation/datasource"
)

func TransformPrefs(prefs datasource.Prefs) datasource.Prefs {

	var (
		result = map[string]map[string]float64{}
	)

	for name, items := range prefs {
		for key, value := range items {

			if _, ok := result[key]; !ok {
				result[key] = map[string]float64{}
			}

			result[key][name] = value
		}
	}

	return result
}

func CalculateSimilarItems(prefs datasource.Prefs) datasource.Prefs {

	var (
		result = map[string]map[string]float64{}
	)

	for name := range prefs {

		result[name] = map[string]float64{}
		scores := TopMatches(prefs, name, EuclideanDistance)

		for _, score := range scores {
			result[name][score.Name] = score.Score
		}
	}

	return result

}
