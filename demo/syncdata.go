package main

import (
	"fmt"
	"sort"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kobeld/recommendation/datasource"
	"github.com/kobeld/recommendation/filters"
	"github.com/kobeld/recommendation/models"
)

type Order struct {
	ID uint `gorm:"primary_key"`

	OrderCode  string `gorm:"size:100;unique_index:uidx_order_code"`
	CustomerID string `gorm:"size:100;index:idx_customer_id"`

	OrderItems []OrderItem `gorm:"ForeignKey:OrderCode;AssociationForeignKey:OrderCode"`
}

type OrderItem struct {
	ID uint `gorm:"primary_key"`

	OrderCode string `gorm:"size:100;index:idx_order_items_order_code"`

	ArticleCode     string `gorm:"size:100"`
	ProductCode     string `gorm:"size:100"`
	Quantity        uint64
	ProductName     string `gorm:"size:255"`
	ProductImageUrl string `gorm:"size:255"`
}

func main() {

	var (
		dbDialect = "postgres"
		dbParams  = "user=aigle password=123 dbname=aigle_dev sslmode=disable host=localhost port=5433"
	)

	db, err := gorm.Open(dbDialect, dbParams)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&models.Purchase{})

	// Load all orders
	var orders []*Order
	err = db.Preload("OrderItems").Find(&orders).Error
	if err != nil {
		panic(err)
	}

	var (
		// Products that have been purchased
		purchasedProducts   = map[string]bool{}
		purchasedUsers      = map[string]bool{}
		userCentricDataSets = map[string]map[string]float64{}
		itemCentricDataSets = map[string]map[string]float64{}
		itemScoresMatrix    = map[string]map[string]float64{}
	)

	// Looping for setting datasets
	for _, order := range orders {
		if len(order.OrderItems) == 0 {
			continue
		}

		dataSet, ok := userCentricDataSets[order.CustomerID]
		if !ok {
			dataSet = map[string]float64{}
		}

		if len(order.OrderItems) > 3 {
			println(order.CustomerID)
		}

		for _, orderItem := range order.OrderItems {
			dataSet[orderItem.ProductCode] = 1.0
			purchasedProducts[orderItem.ProductCode] = true
		}

		userCentricDataSets[order.CustomerID] = dataSet

		purchasedUsers[order.CustomerID] = true
	}

	itemCentricDataSets = filters.TransformPrefs(userCentricDataSets)
	itemScoresMatrix = calculateItemBasedDataset(itemCentricDataSets, purchasedUsers)

	for key1, value1 := range itemScoresMatrix {
		break
		fmt.Printf("\n%s: \n", key1)
		for key2, value2 := range value1 {
			fmt.Printf("    %s: %+v \n", key2, value2)
		}
	}

	// Recommend prodcuts for a user
	targetPerson := "65610"
	userRates := userCentricDataSets[targetPerson]

	// Set all products
	for product, _ := range purchasedProducts {
		if _, ok := userRates[product]; !ok {
			userRates[product] = 0.0
		}
	}

	itemScores := filters.GetRecommendationItems(itemScoresMatrix, userRates)
	for i, ps := range itemScores {
		fmt.Printf("  %d) %s: %+v\n", i+1, ps.Name, ps.Score)
	}

}

func calculateItemBasedDataset(prefs datasource.Prefs, purchasedUsers map[string]bool) datasource.Prefs {

	var (
		result = map[string]map[string]float64{}
	)

	for name := range prefs {

		result[name] = map[string]float64{}
		scores := topMatches(prefs, name, purchasedUsers)

		for _, score := range scores {
			result[name][score.Name] = score.Score
		}
	}

	return result
}

func topMatches(prefs datasource.Prefs, person string, productLists map[string]bool) (ps []filters.ItemScore) {

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

		for productCode := range productLists {
			if _, ok := personItems[productCode]; !ok {
				personItems[productCode] = 0.0
			}

			if _, ok := otherItems[productCode]; !ok {
				otherItems[productCode] = 0.0
			}
		}

		itemScore := filters.ItemScore{
			Name:  name,
			Score: filters.PearsonCorrelation(personItems, otherItems),
		}

		ps = append(ps, itemScore)
	}

	sort.Sort(filters.ByScore(ps))
	if len(ps) > 8 {
		ps = ps[:8]
	}

	return
}
