package models

import (
	"time"

	"github.com/go-faker/faker/v4"
)

type User struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GenerateUser(value int) []User {
	var users []User
	var timeStampLayout = "2006-01-02T15:04:05"

	for i := 0; i < value; i++ {
		timestamp := faker.Timestamp()
		createdTime, err := time.Parse(timeStampLayout, timestamp)
		if err != nil {
			createdTime = time.Now()
		}

		randomWeeks := time.Duration(1 + i%8) // Using modulo to get values between 1-8
		updatedTime := createdTime.Add(time.Hour * 24 * 7 * randomWeeks)

		users = append(users, User{
			ID:        faker.UUIDHyphenated(),
			FirstName: faker.FirstName(),
			LastName:  faker.LastName(),
			UserName:  faker.Username(),
			Email:     faker.Email(),
			CreatedAt: createdTime,
			UpdatedAt: updatedTime,
		})
	}

	return users
}
