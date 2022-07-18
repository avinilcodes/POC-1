package user

type AddUserRequest struct {
	Name     string
	Email    string
	Password string
	RoleType string `json:"role_type"`
}
