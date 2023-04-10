package model

type Meal struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	Occasion int    `json:"occasion"` // 0: Morning, 1: Lunch, 2: Dinner, 3: Snack
	Name     string `json:"name"`
	Date     string `json:"date"`
}
