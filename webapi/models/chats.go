package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	Username     string    `json:"username"`
	CustomName   string    `json:"name"`
	Bio          string    `json:"bio"`
	PasswordHash string    `json:"password"`
	LastOnline   time.Time `json:"time_stamp"`
	Image        []byte    `json:"image"`
}
type UserLogin struct {
	Email        string `json:"email"`
	PasswordHash string `json:"password"`
}

// PROSTI MENYA!

type Chat struct {
	Message     string    `json:"message"`
	MessageDate time.Time `json:"message_date"`
	Name        string    `json:"chat_name"`
	Users       []*User   `json:"users"`
}
type CreateChat struct {
	Name string `json:"chat_name"`
	//will think of it later потому что я таво рот шатала!!! надо придумать как умно получать юзер 1 и юзер 2!!!
	//у себя я через jwt токен который хранит идентити получаю реквестера, но пока не придумал как добавляемого извлечь
}
