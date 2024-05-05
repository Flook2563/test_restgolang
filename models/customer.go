package models

type Customer struct {
	ID   int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
