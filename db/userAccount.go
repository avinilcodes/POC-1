package db

import (
	"context"
	"database/sql"
	"taskmanager/utils"
	"time"

	"github.com/sethvargo/go-password/password"

	"golang.org/x/crypto/bcrypt"
)

const (
	createUserQuery = `INSERT INTO users (
		id,name,email,password,role_type,created_at,updated_at)
		VALUES($1,$2,$3,$4,$5,$6,$7)
	`

	createAccount = `INSERT INTO account(
		id,opening_date,user_id,created_at)
		VALUES($1,$2,$3,$4)`
)

type CreateUserResponse struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//generate random password
func generatePassword() (string, error) {
	return password.Generate(8, 2, 1, false, false)
}

func (s *store) CreateUserAccount(ctx context.Context, user *User) (resp CreateUserResponse, err error) {
	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return
	}
	// defer tx.Rollback()

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}

		err = tx.Commit()
		if err != nil {
			return
		}

	}()

	now := time.Now()
	password, err := generatePassword()
	if err != nil {
		return
	}

	resp = CreateUserResponse{Email: user.Email, Password: password}
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return
	}

	id := utils.GetUniqueId()
	accountId := utils.GetUniqueId()

	_, err = tx.ExecContext(ctx, createUserQuery, id, user.Name, user.Email, hashedPassword, user.RoleType, now, now)
	if err != nil {
		return
	}

	_, err = tx.ExecContext(ctx, createAccount, accountId, now, id, now)
	if err != nil {
		return
	}

	return

}
