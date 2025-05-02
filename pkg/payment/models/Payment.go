package models

import "github.com/google/uuid"

type Payment struct {
	ID     uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key;column:id"`
	Amount float64   `json:"amount" validate:"required" gorm:"column:amount"`
}
