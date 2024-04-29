package httpclient

import "time"

var defaultHeaders = map[string]string{
	"Accept":       "application/json",
	"Content-Type": "application/json",
}

var clientTimeout = time.Second * 10

var debug = true
