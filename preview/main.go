package main

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/badoux/goscraper"
)

type Preview struct {
    Url   	  string   `json:"url"`
    Title         string   `json:"title"`
    SiteName      string   `json:"siteName"`
    Description   string   `json:"description"`
    Images      []string   `json:"images"`
    Image         string   `json:"image"`
}

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse
type Request events.APIGatewayProxyRequest

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, request Request) (Response, error) {
	body, err := url.ParseQuery(request.Body)
	scrapUrl := ""
	if err != nil {
		return Response{StatusCode: 404}, err
    } else {
      scrapUrl = body["url"][0]
    }

	if len(scrapUrl) == 0 {
		return Response{StatusCode: 404}, nil
	}

	s, err := goscraper.Scrape(scrapUrl, 5)
	if err != nil {
	    return Response{StatusCode: 404}, err
	}
	var pvw Preview
	pvw.Url = s.Preview.Link
	pvw.Title = s.Preview.Title
	pvw.Description = s.Preview.Description
	pvw.Images = s.Preview.Images
	pvw.Image = s.Preview.Images[0]
	pvw.SiteName = s.Preview.Name

	result, err := json.Marshal(pvw)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            string(result),
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Access-Control-Allow-Origin": "*",
			"Access-Control-Allow-Credentials": "true",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
