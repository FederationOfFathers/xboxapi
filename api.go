package xboxapi

import "net/http"

var DefaultConfig = &Config{
	APIKey:   "unconfigured",
	Language: "en-US",
}

type Config struct {
	APIKey   string
	Language string
}

// Client allows your to access the xboxapi api
type Client struct {
	*http.Client
}

func New(c *Config) *Client {
	if c == nil {
		c = DefaultConfig
	}
	var rval = &Client{
		Client: &http.Client{
			Transport: &rt{
				cfg:    c,
				client: http.DefaultClient,
			},
		},
	}
	return rval
}
