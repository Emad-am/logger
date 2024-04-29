package httpclient

import (
	"github.com/go-resty/resty/v2"
)

type HttpClient struct {
	Client *resty.Client
}

func NewHttpClient() *HttpClient {

	client := resty.New()

	client = client.SetTimeout(clientTimeout)

	client.SetDebug(debug)

	return &HttpClient{
		Client: client,
	}
}
