package entity

import (
	"errors"
	"time"
)

type User struct {
	ID           uint64     `gorm:"primary_key;auto_increment" json:"id"`
	FirstName    string     `gorm:"size:100;not null;" json:"first_name"`
	LastName     string     `gorm:"size:100;not null;" json:"last_name"`
	Email        string     `gorm:"size:100;not null;unique" json:"email"`
	HashPassword string     `gorm:"size:100;not null;" json:"hash_password"`
	CreatedAt    time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
}

type AuthUserDTO struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserDTO struct {
	ID        ID     `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func (dto *AuthUserDTO) FromDTO() User {

	return User{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
	}
}

func (dto *AuthUserDTO) Validate() error {
	if dto.FirstName == "" {
		return errors.New("first_name is required")
	}
	if dto.LastName == "" {
		return errors.New("last_name is required")
	}
	if dto.Email == "" {
		return errors.New("email is required")
	}
	if dto.Password == "" {
		return errors.New("password is required")
	}

}

func (t *User) ToDTO() UserDTO {
	return UserDTO{
		ID:        ToID(t.ID),
		FirstName: t.FirstName,
		LastName:  t.LastName,
		Email:     t.Email,
	}
}
