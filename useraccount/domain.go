package useraccount

type createRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	RoleType string `json:"role_type"`
}
