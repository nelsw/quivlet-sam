package model

import (
	"encoding/json"
	"github.com/nelsw/quivlet-sam/util/web"
	"strconv"
)

type ChallengeResponse struct {
	Challenges []Challenge `json:"results"`
}

type Challenge struct {
	Category         string   `json:"category"`
	Type             string   `json:"type"`
	Difficulty       string   `json:"difficulty"`
	Question         string   `json:"question"`
	CorrectAnswer    string   `json:"correct_answer"`
	IncorrectAnswers []string `json:"incorrect_answers"`
}

type Problem struct {
	*Token
	*Index
	Category string   `json:"category"`
	Question string   `json:"question"`
	Options  []string `json:"options"`
}

type Solution struct {
	*Token
	*Index
	Answer string `json:"answer"`
}

// Index identifies the challenge number given to Quivlet participants, beginning with 1.
type Index struct {
	Value int `json:"index"`
}

func NewIndex(value string) *Index {
	index := new(Index)
	intValue, _ := strconv.Atoi(value)
	index.Value = intValue
	return index
}

func NewChallenge(token string) Challenge {
	r := &ChallengeResponse{}
	_ = json.Unmarshal(web.Get("https://opentdb.com/api.php?amount=1&type=multiple&encode=base64&token="+token), &r)
	return r.Challenges[0]
}
