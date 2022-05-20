package model

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/Binaretech/classroom-main/storage"
	"gorm.io/gorm"
)

// User model for the database table users
type User struct {
	ID           string `gorm:"primaryKey;type=varchar(32);not null"`
	Name         string `gorm:"type:varchar(64);not null" json:"name"`
	Lastname     string `gorm:"type:varchar(64);not null" json:"lastname"`
	ProfileImage *File  `gorm:"-" json:"profileImage,omitempty"`
	Files        []File `gorm:"polymorphic:Fileable;" json:"files,omitempty"`
	Timestamps
	SoftDeletable
}

func (user *User) UpdateProfileImage(db *gorm.DB, file *multipart.FileHeader, ID string) error {
	if file == nil {
		return nil
	}

	buffer, _ := file.Open()
	defer buffer.Close()

	bytes := bytes.NewBuffer(nil)
	io.Copy(bytes, buffer)

	mimeType := file.Header.Get("Content-Type")

	key := fmt.Sprintf("%s/profile-image", ID)
	if err := storage.Put(storage.USERS_BUCKET, key, bytes.Bytes(), mimeType); err != nil {
		return err
	}

	profileImage := File{
		Key:      key,
		MimeType: mimeType,
		Bucket:   storage.USERS_BUCKET,
	}

	if err := db.Model(&User{ID: user.ID}).Association("Files").Replace(&profileImage); err != nil {
		return err
	}

	user.ProfileImage = &profileImage

	return nil

}

// StoreUserRequest is the request to store a user in the database
type StoreUserRequest struct {
	ID       string `form:"id" json:"id" validate:"required,unique=users;id"`
	Name     string `form:"name" json:"name" validate:"required,max=64"`
	Lastname string `form:"lastname" json:"lastname" validate:"required,max=64"`
}

// User return a user model with the request data
func (req *StoreUserRequest) User() User {
	return User{
		ID:       req.ID,
		Name:     req.Name,
		Lastname: req.Lastname,
	}
}

// UpdateUserRequest is the request to update a user in the database
type UpdateUserRequest struct {
	ID       string `json:"id" validate:"required,exists=users;id"`
	Name     string `json:"name" validate:"omitempty,max=64"`
	Lastname string `json:"lastname" validate:"omitempty,max=64"`
}

// Data returns data to update a user in the database
func (req *UpdateUserRequest) Data() map[string]interface{} {
	return map[string]interface{}{
		"name":     req.Name,
		"lastname": req.Lastname,
	}
}
