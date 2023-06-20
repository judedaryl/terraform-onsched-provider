package onsched

import (
	"context"
	"fmt"
	"net/http"

	oauth2 "golang.org/x/oauth2/clientcredentials"
)

type Client struct {
	http    *http.Client
	env     Environment
	apiHost string
}

type Environment int64

const (
	Sandbox Environment = iota
	Prod
)

func hostBuilder(service string, env Environment) string {
	baseUrl := "onsched.com"
	if env == Sandbox {
		baseUrl = fmt.Sprintf("https://%s-%s.%s", "sandbox", service, baseUrl)
	}
	return baseUrl
}

func apiHost(env Environment) string {
	return hostBuilder("api", env)
}
func identityHost(env Environment) string {
	return hostBuilder("identity", env)
}

func NewClient(env Environment, client_id, client_secret string) *Client {
	return NewClientWithContext(env, client_id, client_secret, context.Background())
}

func NewClientWithContext(env Environment, client_id, client_secret string, ctx context.Context) *Client {
	url := identityHost(env)
	conf := &oauth2.Config{
		ClientID:     client_id,
		ClientSecret: client_secret,
		Scopes:       []string{"OnSchedApi"},
		TokenURL:     fmt.Sprintf("%s/connect/token", url),
	}
	return &Client{
		http:    conf.Client(ctx),
		env:     env,
		apiHost: apiHost(env),
	}
}

func (c *Client) GetCompany() (Company, error) {
	result, err := c.get("setup/v1/companies")
	if err != nil {
		return Company{}, err
	}
	return parse[Company](result)
}

func (c *Client) UpdateCompany(company Company) (Company, error) {
	result, err := c.put("setup/v1/companies", company)
	if err != nil {
		return Company{}, err
	}
	return parse[Company](result)
}
