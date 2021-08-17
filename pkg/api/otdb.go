package api

const OpenTDBTokenRequestUrl = "https://opentdb.com/api_token.php?command=request"

type Code int

const (
	Success    = Code(iota) // Success, Returned results successfully.
	NoResults               // NoResults, Could not return results. API doesn't have enough questions for your query.
	BadParam                // BadParam, Contains an invalid parameter. Arguments passed in aren't valid.
	TokenUnk                // TokenUnk, Session Token does not exist.
	TokenEmpty              // TokenEmpty, Session Token has returned all questions for the specified query. Must reset.
)

type TokenRequestResponse struct {
	ResponseCode    int    `json:"response_code"`
	ResponseMessage string `json:"response_message"`
	Token           string `json:"token"`
}
