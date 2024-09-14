package utils

type FoodItem struct {
	Food  string
	Price float64
}

func InitData() []FoodItem {
	foods := []FoodItem{
		{Food: "Pizza", Price: 10.50},
		{Food: "Pasta", Price: 8.00},
		{Food: "Salad", Price: 5.50},
	}
	return foods
}
