package data

import "health-app/server/model"

var InitUsers = []model.User{
	{
		ID:        0,
		Name:      "Admin",
		Password:  "password",
		Birthday:  "1990-01-01",
		Meals:     []model.Meal{},
		Records:   []model.Record{},
		Exercises: []model.Exercise{},
		Diary:     []model.Diary{},
	},
	{
		ID:        1,
		Name:      "Moderator",
		Password:  "password",
		Birthday:  "1990-01-01",
		Meals:     []model.Meal{},
		Records:   []model.Record{},
		Exercises: []model.Exercise{},
		Diary:     []model.Diary{},
	},
}
