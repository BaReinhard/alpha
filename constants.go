package alpha

import (
	"net/http"
)

// APIURL the base API url
const APIURL = "https://www.alphavantage.co/query"

// Options are the options passed to create a client
type Options struct {
	Function   string `json:"function"`
	Symbol     string `json:"symbol"`
	OutputSize string `json:"outputsize"`
	DataType   string `json:"datatype"`
	APIKey     string `json:"apikey"`
}

// Client is the struct returned and used for API calls
type Client struct {
	Options Options
	Client  *http.Client
	base    string
}

// NewClient returns a new Alpha API Client
func NewClient(opts Options, client *http.Client) Client {
	c := Client{}
	c.base = APIURL
	c.Options = opts
	c.Client = client
	return c
}
