package data

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User 유형에 대한 데이터를 설명합니다.
type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	IsAdmin   int       `json:"is_admin"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// PasswordMatches는 Go의 bcrypt 패키지를 사용하여 사용자가 제공한 비밀번호와
// 데이터베이스에 특정 사용자에 대해 저장한 해시를 비교합니다. 비밀번호와 해시가 일치하면
// 참을 반환하고, 그렇지 않으면 거짓을 반환합니다.
func (u *User) PasswordMatches(plainText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainText))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			// invalid password
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}
