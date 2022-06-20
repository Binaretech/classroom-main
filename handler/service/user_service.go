package service

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/Binaretech/classroom-main/db/model"
	"github.com/Binaretech/classroom-main/storage"
	"gorm.io/gorm"
)

func UpdateProfileImage(user model.User, db *gorm.DB, file *multipart.FileHeader, ID string) error {
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

	profileImage := model.File{
		Key:      key,
		MimeType: mimeType,
		Bucket:   storage.USERS_BUCKET,
	}

	if err := db.Model(&model.User{ID: user.ID}).Association("Files").Replace(&profileImage); err != nil {
		return err
	}

	user.ProfileImage = &profileImage

	return nil

}
