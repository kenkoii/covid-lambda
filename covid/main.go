package main

import (
	"bytes"
	"encoding/json"

	"github.com/ahmednafies/go-covid/covid"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(req events.APIGatewayProxyRequest) (Response, error) {
	var buf bytes.Buffer
	country := req.QueryStringParameters["country"]

	var res interface{}
	if country != "" {
		data := covid.GetCountryByName(country)
		res = Case{
			Country: Country{
				Name:      data.Attrs.Country,
				Latitude:  data.Attrs.Latitude,
				Longitude: data.Attrs.Longitude,
			},
			Confirmed:  data.Attrs.Confirmed,
			Deaths:     data.Attrs.Deaths,
			Active:     data.Attrs.Active,
			Recovered:  data.Attrs.Recovered,
			LastUpdate: data.Attrs.LastUpdate / 1000,
		}
	} else {
		cases := []Case{}
		data := covid.GetData()
		for _, v := range data {
			cases = append(cases, Case{
				Country: Country{
					Name:      v.Attrs.Country,
					Latitude:  v.Attrs.Latitude,
					Longitude: v.Attrs.Longitude,
				},
				Confirmed:  v.Attrs.Confirmed,
				Deaths:     v.Attrs.Deaths,
				Active:     v.Attrs.Active,
				Recovered:  v.Attrs.Recovered,
				LastUpdate: v.Attrs.LastUpdate / 1000,
			})
			res = cases
		}
	}

	body, err := json.Marshal(res)
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
