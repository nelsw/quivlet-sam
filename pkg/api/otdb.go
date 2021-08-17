package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type TokenResponse struct {
	Code    `json:"response_code"`
	Message string `json:"response_message"`
	Token   string `json:"token"`
}

type ProblemResponse struct {
	Code     `json:"response_code"`
	Problems []Problem `json:"results"`
}

type Problem struct {
	Category         string   `json:"category"`
	Type             string   `json:"type"`
	Difficulty       string   `json:"difficulty"`
	Question         string   `json:"question"`
	CorrectAnswer    string   `json:"correct_answer"`
	IncorrectAnswers []string `json:"incorrect_answers"`
}

type Code int

const (
	Success    = Code(iota) // Success, Returned results successfully.
	NoResults               // NoResults, Could not return results. API doesn't have enough questions for your query.
	BadParam                // BadParam, Contains an invalid parameter. Arguments passed in aren't valid.
	TokenUnk                // TokenUnk, Session Token does not exist.
	TokenEmpty              // TokenEmpty, Session Token has returned all questions for the specified query. Must reset.

	tokenUrl   = "https://opentdb.com/api_token.php?command=request"
	problemUrl = "https://opentdb.com/api.php?amount=1&type=multiple&encode=base64"
)

var client = &http.Client{Timeout: 25 * time.Second}

func NewToken() string {
	r := &TokenResponse{}
	_ = json.Unmarshal(get(tokenUrl), &r)
	return r.Token
}

func NewProblem(token string) Problem {
	r := &ProblemResponse{}
	_ = json.Unmarshal(get(problemUrl+"&token="+token), &r)
	return r.Problems[0]
}

func get(url string) []byte {
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	response, _ := client.Do(request)
	body, _ := ioutil.ReadAll(response.Body)
	_ = response.Body.Close()
	return body
}
