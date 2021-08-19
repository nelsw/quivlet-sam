package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/nelsw/quivlet-sam/model"
	"github.com/nelsw/quivlet-sam/util/api"
	"github.com/nelsw/quivlet-sam/util/random"
	"github.com/nelsw/quivlet-sam/util/transform"
)

// HandleRequest is responsible for receiving requests regarding answer submission and problem generation.
func HandleRequest(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	api.Log(r)

	c := model.Challenge{}
	transform.UnmarshalStr(r.Body, &c)

	// if this path variable is present, a user is submitting an answer
	solve := r.QueryStringParameters["solve"]
	if solve != "" {
		c.FindChallenge()
		return api.Response(200, &c)
	}

	token := c.Token
	index := c.Index
	c.FindChallenge()
	if c.Question != "" {
		c.IncorrectAnswers = append(c.IncorrectAnswers, c.CorrectAnswer)
		random.Shuffle(c.IncorrectAnswers)
		c.CorrectAnswer = ""
		return api.Response(200, &c)
	}

	c = model.NewChallenge(token)
	c.Token = token
	c.Index = index
	c.Category = transform.FromBase64(c.Category)
	c.Type = transform.FromBase64(c.Type)
	c.Difficulty = transform.FromBase64(c.Difficulty)
	c.Question = transform.FromBase64(c.Question)
	c.CorrectAnswer = transform.FromBase64(c.CorrectAnswer)
	for i, answer := range c.IncorrectAnswers {
		c.IncorrectAnswers[i] = transform.FromBase64(answer)
	}

	c.SaveChallenge()
	c.IncorrectAnswers = append(c.IncorrectAnswers, c.CorrectAnswer)
	random.Shuffle(c.IncorrectAnswers)
	c.CorrectAnswer = ""

	return api.Response(200, &c)

}

func main() {
	lambda.Start(HandleRequest)
}
