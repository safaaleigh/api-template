package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

// Todo a single todo
type Todo struct {
	ID         uuid.UUID  `json:"id"`
	CreateDate time.Time  `json:"create_date" db:"create_date"`
	UpdateDate time.Time  `json:"update_date" db:"update_date"`
	Name       string     `json:"name" valid:"required"`
	Completed  bool       `json:"completed"`
	Due        *time.Time `json:"due"`
}

// Todos a list of todos
type Todos []Todo

// Validate validates the Todo fields
func (m Todo) Validate() error {
	_, err := govalidator.ValidateStruct(m)
	if err != nil {
		return err
	}
	return nil
}
