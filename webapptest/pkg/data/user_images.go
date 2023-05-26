package data

import "time"

// UserImage는 사용자 프로필 이미지의 유형입니다.
type UserImage struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	FileName  string    `json:"file_name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
