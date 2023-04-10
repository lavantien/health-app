package model

type User struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"` // unique
	Password  string     `json:"password"`
	Birthday  string     `json:"birthday"`
	Meals     []Meal     `json:"meals"`
	Records   []Record   `json:"records"`
	Exercises []Exercise `json:"exercises"`
	Diary     []Diary    `json:"diary"`
}
