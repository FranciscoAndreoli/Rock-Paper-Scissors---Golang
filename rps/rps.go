package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK         = 0 // beats scissors. (scissors + 1) % 3 = 0
	PAPER        = 1 // beats rock. (rock + 1) % 3 = 1
	SCISSORS     = 2 // beats paper. (paper + 1) % 3 = 2
	PLAYERWINS   = 1
	COMPUTERWINS = 2
	DRAW         = 3
)

type Round struct {
	Winner         int    `json:"winner"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
}

func PlayRound(playerValue int) Round {

	ganador := [4]string{"Bien hecho, ganaste!", "Has ganado!", "Excelente!", "Te felicito"}
	perdedor := [4]string{"Eres un perdedor!", "Mejor suerte la próxima!", "Dedicate a otra cosa!", "Intentalo otra vez"}

	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)
	computerChoice := ""
	roundResult := ""
	winner := 0

	switch computerValue {
	case ROCK:
		computerChoice = "La computadora eligió piedra"
		break
	case PAPER:
		computerChoice = "La computadora eligió papel"
		break
	case SCISSORS:
		computerChoice = "La computadora eligió tijera"
		break
	default:

	}

	if playerValue == computerValue {
		roundResult = "Es un empate!"
		winner = DRAW
	} else if playerValue == (computerValue+1)%3 {
		roundResult = ganador[rand.Intn(3)]
		winner = PLAYERWINS
	} else {
		roundResult = perdedor[rand.Intn(3)]
		winner = COMPUTERWINS
	}

	var result Round
	result.Winner = winner
	result.ComputerChoice = computerChoice
	result.RoundResult = roundResult
	return result
}
