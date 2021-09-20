package models

import "time"

type User struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"password"`
	IsActicve bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (user *User) TableName() string {
	return "user"
}

type UserLogin struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}
type UserRegister struct {
	Email     string `form:"email" binding:"required" json:"email"`
	Password  string `form:"password" binding:"required" json:"password"`
	FirstName string `form:"first_name"`
	LastName  string `form:"last_name"`
}

func (user *User) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = user.ID
	resp["email"] = user.Email
	resp["first_name"] = user.FirstName
	resp["last_name"] = user.LastName
	resp["is_active"] = user.IsActicve
	resp["created_at"] = user.CreatedAt
	resp["updated_at"] = user.UpdatedAt
	return resp

}
