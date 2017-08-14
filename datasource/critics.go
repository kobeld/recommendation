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

var BoughtItems Prefs = map[string]map[string]float64{
	"Lisa": map[string]float64{
		"Lady in the Water":  1.0,
		"Snakes on a Plane":  1.0,
		"Just My Luck":       0.0,
		"Superman Returns":   0.0,
		"The Night Listener": 0.0,
		"You, Me and Dupree": 0.0},
	"Gene": map[string]float64{
		"Lady in the Water":  0.0,
		"Snakes on a Plane":  0.0,
		"Just My Luck":       1.0,
		"Superman Returns":   0.0,
		"The Night Listener": 0.0,
		"You, Me and Dupree": 0.0},
	"Michael": map[string]float64{
		"Lady in the Water":  1.0,
		"Snakes on a Plane":  1.0,
		"Just My Luck":       0.0,
		"Superman Returns":   1.0,
		"The Night Listener": 1.0,
		"You, Me and Dupree": 0.0},
	"Claudia": map[string]float64{
		"Lady in the Water":  0.0,
		"Snakes on a Plane":  0.0,
		"Just My Luck":       1.0,
		"Superman Returns":   0.0,
		"The Night Listener": 0.0,
		"You, Me and Dupree": 1.0},
	"Mick": map[string]float64{
		"Lady in the Water":  1.0,
		"Snakes on a Plane":  0.0,
		"Just My Luck":       1.0,
		"Superman Returns":   1.0,
		"The Night Listener": 1.0,
		"You, Me and Dupree": 0.0},
	"Jack": map[string]float64{
		"Lady in the Water":  0.0,
		"Snakes on a Plane":  0.0,
		"Just My Luck":       0.0,
		"Superman Returns":   0.0,
		"The Night Listener": 1.0,
		"You, Me and Dupree": 0.0},
	"Toby": map[string]float64{
		"Lady in the Water":  0.0,
		"Snakes on a Plane":  1.0,
		"Just My Luck":       0.0,
		"Superman Returns":   1.0,
		"The Night Listener": 1.0,
		"You, Me and Dupree": 0.0},
}

var BoughtBooks Prefs = map[string]map[string]float64{
	"1": map[string]float64{
		"Book A": 1.0,
		"Book B": 1.0,
		"Book C": 0.0,
		"Book D": 0.0},
	"2": map[string]float64{
		"Book A": 0.0,
		"Book B": 1.0,
		"Book C": 0.0,
		"Book D": 0.0},
	"3": map[string]float64{
		"Book A": 1.0,
		"Book B": 1.0,
		"Book C": 0.0,
		"Book D": 0.0},
	"4": map[string]float64{
		"Book A": 1.0,
		"Book B": 0.0,
		"Book C": 0.0,
		"Book D": 1.0},
	"5": map[string]float64{
		"Book A": 0.0,
		"Book B": 0.0,
		"Book C": 0.0,
		"Book D": 0.0},
	"6": map[string]float64{
		"Book A": 1.0,
		"Book B": 1.0,
		"Book C": 0.0,
		"Book D": 0.0},
	"7": map[string]float64{
		"Book A": 1.0,
		"Book B": 0.0,
		"Book C": 1.0,
		"Book D": 0.0},
	"8": map[string]float64{
		"Book A": 0.0,
		"Book B": 0.0,
		"Book C": 1.0,
		"Book D": 0.0},
	"9": map[string]float64{
		"Book A": 1.0,
		"Book B": 1.0,
		"Book C": 0.0,
		"Book D": 0.0},
	"10": map[string]float64{
		"Book A": 1.0,
		"Book B": 1.0,
		"Book C": 0.0,
		"Book D": 0.0},
	"11": map[string]float64{
		"Book A": 0.0,
		"Book B": 1.0,
		"Book C": 0.0,
		"Book D": 0.0},
	"12": map[string]float64{
		"Book A": 0.0,
		"Book B": 1.0,
		"Book C": 0.0,
		"Book D": 0.0},
	"13": map[string]float64{
		"Book A": 1.0,
		"Book B": 1.0,
		"Book C": 0.0,
		"Book D": 0.0},
	"14": map[string]float64{
		"Book A": 0.0,
		"Book B": 0.0,
		"Book C": 0.0,
		"Book D": 0.0},
	"15": map[string]float64{
		"Book A": 0.0,
		"Book B": 0.0,
		"Book C": 0.0,
		"Book D": 0.0},
}
