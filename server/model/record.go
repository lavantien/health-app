package model

type Record struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Date    string `json:"date"`
	Weight  int    `json:"weight"`
	BodyFat int    `json:"body_fat"`
}
