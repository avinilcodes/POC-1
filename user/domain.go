package user

type updateRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
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
