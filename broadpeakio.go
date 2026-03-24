package broadpeakio

import (
	"fmt"
)

type BroadpeakClient struct {
	apiKey   string
	Endpoint string
}

type Identifiable struct {
	Id uint `json:"id,omitempty"`
}

type ApiResponseOutput struct {
	StatusCode uint   `json:"statusCode"`
	Message    string `json:"message"`
	Error      string `json:"error"`
}

func MakeClient(apiKey string, endpoint string) BroadpeakClient {
	// default value if it is not set
	if endpoint == "" {
		endpoint = "https://api.broadpeak.io"
	}
	client := BroadpeakClient{
		apiKey:   apiKey,
		Endpoint: endpoint,
	}
	return client
}

func (c *BroadpeakClient) getBaseUrl() string {
	return fmt.Sprintf("%s/v1/", c.Endpoint)
}
