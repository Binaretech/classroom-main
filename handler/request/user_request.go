package request

import "github.com/Binaretech/classroom-main/db/model"

// StoreUserRequest is the request to store a user in the database
type StoreUserRequest struct {
	ID       string `form:"id" json:"id" validate:"required,unique=users;id"`
	Name     string `form:"name" json:"name" validate:"required,max=64"`
	Lastname string `form:"lastname" json:"lastname" validate:"required,max=64"`
}

// User return a user model with the request data
func (req *StoreUserRequest) User() model.User {
	return model.User{
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
