package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// const postmarkServerToken = "9879d045-32a9-4deb-a159-6e64c9e56539" // vespene
const postmarkServerToken = "aa047268-111d-4a4b-be4d-552ad62e3cb0" // boxd

func main() {
	// s := `{"Message":"Full Name: Luis Vega\nEmail: eldosoa@gmail.com\nMessage:\n\nI like this message!"}`

	// b := body{}

	// if err := json.Unmarshal([]byte(s), &b); err != nil {
	// 	panic(err)
	// }

	// if err := send(b); err != nil {
	// 	panic(err)
	// }

	// return

	lambda.Start(handler)
}

type body struct {
	Message     string
	Attachments []struct {
		Name        string
		Content     string
		ContentType string
	}
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	b := body{}

	if err := json.Unmarshal([]byte(request.Body), &b); err != nil {
		return nil, err
	}

	if err := send(b); err != nil {
		return nil, err
	}

	return &events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

func send(b body) error {
	outload := map[string]interface{}{
		"From":        "BOxD Website <website@boxd.us>",
		"To":          "hello@boxd.us",
		"Subject":     "[BOxD Website] We received a message!",
		"TextBody":    b.Message,
		"Attachments": b.Attachments,
	}

	var buffer bytes.Buffer

	if err := json.NewEncoder(&buffer).Encode(outload); err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "https://api.postmarkapp.com/email", &buffer)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Postmark-Server-Token", postmarkServerToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		bs, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("send failed: %s", string(bs))
	}

	return nil
}
