package user

type updateRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type AddUserRequest struct {
	Name     string
	Email    string
	Password string
	RoleType string `json:"role_type"`
}

func (ur updateRequest) Validate() (err error) {
	if ur.Password == "" {
		return errEmptyPassword
	}
	if ur.Name == "" {
		return errEmptyName
	}
	return
}
