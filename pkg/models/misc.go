package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Food_Item struct {
	UserId   uint    `json:"user_id"`
	Name     string  `json:"name"`
	Kcal     float32 `json:"kcal"`
	Protein  float32 `json:"protein"`
	Fiber    float32 `json:"fiber"`
	Fat      float32 `json:"fat"`
	Carbs    float32 `json:"carbs"`
	Standard uint    `json:"standard"`

	// can add more macro nutrients later
	gorm.Model
}

type FoodConsumption struct {
	UserID       uint           `json:"user_id"`
	Food_Item_ID uint           `json:"food_item_id"`
	Quantity     float32        `json:"quantity"`
	Date         datatypes.Date `json:"date"`
	ID           uint           `gorm:"primarykey"`
}

type UserFoodRequirements struct {
	UserID  uint    `json:"user_id"`
	Kcal    float32 `json:"kcal"`
	Protein float32 `json:"protein"`
	Fiber   float32 `json:"fiber"`
	Fat     float32 `json:"fat"`
	Carbs   float32 `json:"carbs"`
	Date    string  `json:"date"`
}

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  uint   `json:"pages"`
	Status uint   `json:"status"` // read, reading, to read
	gorm.Model
}
