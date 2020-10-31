package data

import "time"

type Thread struct {
	Uuid          string
	User          User
	Numreplies    int
	Topic         string
	CreatedAtDate time.Time
}

func Threads() ([]Thread, error) {
	return []Thread{
		{
			Uuid: "Uuid001",
			User: User{
				Name: "Apple",
			},
			Numreplies:    2,
			Topic:         "Buy Apple",
			CreatedAtDate: time.Now(),
		},
		{
			Uuid: "Uuid002",
			User: User{
				Name: "Banana",
			},
			Numreplies:    4,
			Topic:         "Buy Banana",
			CreatedAtDate: time.Now(),
		},
		{
			Uuid: "Uuid003",
			User: User{
				Name: "Orange",
			},
			Numreplies:    0,
			Topic:         "Buy Orange",
			CreatedAtDate: time.Now(),
		},
	}, nil
}
