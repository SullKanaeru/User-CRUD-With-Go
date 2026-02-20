package model

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"` 
	Name      string    `json:"name"`                 // Pakai Name, bukan Fullname
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}