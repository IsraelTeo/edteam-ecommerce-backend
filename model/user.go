package model

import (
	"encoding/json"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID       `json:"id"`
	Email     string          `json:"email"`
	Password  string          `json:"password"`
	IsAdmin   bool            `json:"is_admin"`
	Details   json.RawMessage `json:"details"`    // Detalles adicionales del usuario serán un json manipulado por el front-end
	CreatedAt int64           `json:"created_at"` // Fecha de registro del usuario será de tipo int64 para trabajar con horarios UnixTime
	UpdatedAt int64           `json:"updated_at"` // Fecha de actualización del usuario será de tipo int64 para trabajar con horarios UnixTime
}

type Users []User
