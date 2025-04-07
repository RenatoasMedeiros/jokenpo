package models

import "github.com/google/uuid"

type Games struct {
	ID            uuid.UUID `json:id`
	Player1       uuid.UUID `json:player_1`
	Player2       uuid.UUID `json:player_2`
	Player1Choice string    `json:player_1_choice`
	Player2Choice string    `json:player_2_choice`
	Winner        uuid.UUID `json:winner`
}

type GameInput struct {
	Player1       string `json:"player_1"`
	Player2       string `json:"player_2"`
	Player1Choice string `json:"player_1_choice"`
	Player2Choice string `json:"player_2_choice"`
}
