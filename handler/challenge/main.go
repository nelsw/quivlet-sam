package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/nelsw/quivlet-sam/model"
	"github.com/nelsw/quivlet-sam/util/api"
	"github.com/nelsw/quivlet-sam/util/random"
	"github.com/nelsw/quivlet-sam/util/transform"
)

// HandleRequest is responsible for receiving requests regarding answer submission and problem generation.
func HandleRequest(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	api.Log(r)

	// if this path variable is present, a user is submitting an answer
	solve := r.QueryStringParameters["solve"]
	if solve != "" {

	}

	challenge := model.NewChallenge(r.QueryStringParameters["token"])

	solution := &model.Solution{}

	solution.Answer = transform.FromBase64(challenge.CorrectAnswer)
	// save

	for i, answer := range challenge.IncorrectAnswers {
		challenge.IncorrectAnswers[i] = transform.FromBase64(answer)
	}

	problem := &model.Problem{}
	problem.Token = model.NewToken(r.QueryStringParameters["token"])
	problem.Index = model.NewIndex(r.QueryStringParameters["index"])
	problem.Question = transform.FromBase64(challenge.Question)
	problem.Options = challenge.IncorrectAnswers
	problem.Options = append(problem.Options, solution.Answer)

	random.Shuffle(problem.Options)

	return api.Response(200, &problem)

}
