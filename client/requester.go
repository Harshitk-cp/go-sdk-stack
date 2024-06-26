package client

import (
	"net/http"
)

type Requester interface {
    Do(req *http.Request) (*http.Response, error)
}
