package models

type User struct {
	ID       string `gorm:"primaryKey;autoIncrement:false;type:varchar(255);not null;unique" json:"id"`
	Name     string `gorm:"type:varchar(255);not null" json:"name"`
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Active   bool   `gorm:"not null;default:true" json:"is_active"`
	Admin    bool   `gorm:"not null;default:false" json:"is_admin"`
}

func NewUser(userID, name, email string, isActive, isAdmin bool) *User {
	return &User{
		ID:       userID,
		Name:     name,
		Email:    email,
		Active:   isActive,
		Admin:    isAdmin,
	}
}
