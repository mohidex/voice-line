package models

type User struct {
	ID       string `gorm:"primaryKey;autoIncrement:false;type:varchar(255);not null;unique" json:"id"`
	Name     string `gorm:"type:varchar(255);not null" json:"name"`
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	PhotoURL string `gorm:"size:255" json:"photo_url"`
	Active   bool   `gorm:"not null;default:true" json:"is_active"`
	Admin    bool   `gorm:"not null;default:false" json:"is_admin"`
}

func NewUser(userID, name, username, email, photoURL string, isActive, isAdmin bool) *User {
	return &User{
		ID:       userID,
		Name:     name,
		Username: username,
		Email:    email,
		PhotoURL: photoURL,
		Active:   isActive,
		Admin:    isAdmin,
	}
}
