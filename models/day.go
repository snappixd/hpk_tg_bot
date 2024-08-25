package models

type Day struct {
	Date    string `json:"date"`
	DayName string `json:"dayName"`
	Paras   []Para `json:"paras"`
}
