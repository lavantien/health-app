package model

type Diary struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Date    string `json:"date"`
	Time    string `json:"time"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
