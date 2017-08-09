package datasource

type Prefs map[string]map[string]float64

var Critics Prefs = map[string]map[string]float64{
	"Lisa": map[string]float64{
		"Lady in the Water":  2.5,
		"Snakes on a Plane":  3.5,
		"Just My Luck":       3.0,
		"Superman Returns":   3.5,
		"You, Me and Dupree": 2.5,
		"The Night Listener": 3.0},
	"Gene": map[string]float64{
		"Lady in the Water":  3.0,
		"Snakes on a Plane":  3.5,
		"Just My Luck":       1.5,
		"Superman Returns":   5.0,
		"You, Me and Dupree": 3.5,
		"The Night Listener": 3.0},
	"Michael": map[string]float64{
		"Lady in the Water":  2.5,
		"Snakes on a Plane":  3.0,
		"Superman Returns":   3.5,
		"The Night Listener": 4.0},
	"Claudia": map[string]float64{
		"Snakes on a Plane":  3.5,
		"Just My Luck":       3.0,
		"The Night Listener": 4.5,
		"Superman Returns":   4.0,
		"You, Me and Dupree": 2.5},
	"Mick": map[string]float64{
		"Lady in the Water":  3.0,
		"Snakes on a Plane":  4.0,
		"Just My Luck":       2.0,
		"Superman Returns":   3.0,
		"The Night Listener": 3.0,
		"You, Me and Dupree": 2.0},
	"Jack": map[string]float64{
		"Lady in the Water":  3.0,
		"Snakes on a Plane":  4.0,
		"The Night Listener": 3.0,
		"Superman Returns":   5.0,
		"You, Me and Dupree": 3.5},
	"Toby": map[string]float64{
		"Snakes on a Plane":  4.5,
		"You, Me and Dupree": 1.0,
		"Superman Returns":   4.0},
}
