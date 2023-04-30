package entities

import "github.com/pborman/uuid"

type Tweet struct {
	ID          string `json:"id"` //renomeando o atributo para visualização em json
	Description string `json:"description"`
}

func NewTwwet() *Tweet {
	tweet := Tweet{
		ID: uuid.New(),
	}

	return &tweet
}
