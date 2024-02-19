package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// Create User
type CreateUserRequest struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Profission string `json:"profission"`
}

func (r *CreateUserRequest) Validate() error {

	if r.Name == "" && r.Email == "" && r.Password == "" && r.Profission == "" {
		return fmt.Errorf("request body is empty or malformed")
	}

	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}

	if r.Email == "" {
		return errParamIsRequired("email", "string")
	}

	if r.Password == "" {
		return errParamIsRequired("password", "string")
	}

	if r.Profission == "" {
		return errParamIsRequired("profission", "string")
	}

	return nil
}

// Update User

type UpdateUserRequest struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Profission string `json:"profission"`
}

func (r *UpdateUserRequest) Validate() error {
	// If any field is provided, validation is truthy
	if r.Name != "" || r.Email != "" || r.Password != "" || r.Profission != "" {
		return nil
	}

	// If none of the fields were provided, return false
	return fmt.Errorf("at least one valid field must be provided")
}
