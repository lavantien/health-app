package model

type Exercise struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	Duration int    `json:"duration"`
	Calories int    `json:"calories"`
	Date     string `json:"date"`
}
