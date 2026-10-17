package main

import "encoding/json"

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MenuItem struct {
	ID              int             `json:"id"`
	Name            string          `json:"name"`
	CategoryID      int             `json:"category_id"`
	CategoryName    string          `json:"category_name"`
	Description     string          `json:"description"`
	Ingredients     string          `json:"ingredients"`
	ImageID         *int            `json:"image_id"`
	ImageURL        string          `json:"image_url"`
	ImageAttributes json.RawMessage `json:"image_attributes"`
	Price           int             `json:"price"`
}

// MenuInput - тело запросов POST/PUT /api/menu с правилами валидации.
type MenuInput struct {
	Name        string `json:"name" binding:"required,max=120"`
	CategoryID  int    `json:"category_id" binding:"required,gt=0"`
	Description string `json:"description" binding:"max=2000"`
	Ingredients string `json:"ingredients" binding:"max=2000"`
	ImageID     *int   `json:"image_id"`
	Price       int    `json:"price" binding:"required,gt=0"`
}

// MenuFilter - параметры выборки списка блюд.
type MenuFilter struct {
	CategoryID int
	Search     string
	SortCol    string
	Dir        string
	Page       int
	Limit      int
}
