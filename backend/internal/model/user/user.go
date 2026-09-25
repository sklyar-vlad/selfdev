package model

import "github.com/google/uuid"

type User struct {
	UserId       uuid.UUID
	Username     string
	Email        string
	PasswordHash string
	AvatarURL    string
}

func NewUser(username, email, passwordHash, avatarURL string) User {
	return User{UserId: uuid.New(), Username: username, Email: email, PasswordHash: passwordHash, AvatarURL: avatarURL}
}
