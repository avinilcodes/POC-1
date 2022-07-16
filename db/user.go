package db

import (
	"context"
	"database/sql"
	"taskmanager/utils"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	createQuery = `INSERT INTO users (
		id,name,email,password,role_type,created_at,updated_at)
		VALUES($1,$2,$3,$4,$5,$6,$7)
	`
	findUserByEmailQuery = `SELECT id, name, email, password, role_type, created_at,updated_at FROM users WHERE email = $1`
	findAllQuery         = `SELECT * FROM users`
	deleteUserByIDQuery  = `DELETE FROM users WHERE id = $1`
	updateUserQuery      = "UPDATE users SET name=$1 ,password=$2,updated_at=$3 where id=$4"
	userInsert           = `INSERT INTO users (id,name,email,password,role_type,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7)`
	//Super User details
	superAdminEmail    = "superadmin@company.com"
	superAdminPassword = "Josh@123"
	superAdminName     = "Josh"
	superAdminRoleType = "super_admin"
)

type User struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	RoleType  string    `json:"role_type" db:"role_type"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s *store) UpdateUser(ctx context.Context, user *User) (err error) {
	now := time.Now()
	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		_, err = s.db.Exec(
			updateUserQuery,
			user.Name,
			user.Password,
			now,
			user.ID,
		)
		return err
	})
}

func (s *store) FindUserByEmail(ctx context.Context, email string) (user User, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.GetContext(ctx, &user, findUserByEmailQuery, email)
	})
	return
}

func CreateSuperAdmin(s *store) (err error) {
	var user User
	err = s.db.QueryRow(findUserByEmailQuery, superAdminEmail).Scan(&user)
	flag := user == User{}
	if flag {
		err = nil
		now := time.Now()
		password, err := HashPassword(superAdminPassword)
		if err != nil {
			return err
		}
		err = nil
		_, err = s.db.Exec(createQuery,
			utils.GetUniqueId(),
			superAdminName,
			superAdminEmail,
			password,
			superAdminRoleType,
			now,
			now,
		)
		if err != nil {
			return err
		}
	}
	return
}

func (s *store) ListUsers(ctx context.Context) (users []User, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.SelectContext(ctx, &users, findAllQuery)
	})
	return
}

func (s *store) CreateUser(ctx context.Context, user User) (err error) {
	password, err := HashPassword(user.Password)
	if err != nil {
		return
	}
	_, err = s.db.Query(userInsert, user.ID, user.Name, user.Email, password, user.RoleType, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}
	return
}
