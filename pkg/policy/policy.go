package policy

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"net/http"
	"os"
)

type Client struct {
	restClient *resty.Client
	config     Config
}

type Config struct {
	Url string
	Key string
}

func newConfigFromEnvironment() (Config, error) {

	policyEndpointUrl, ok := os.LookupEnv("PolicyEndpointUrl")

	if !ok {
		return Config{}, errors.New("not defined PolicyEndpointUrl environment variable")
	}

	policyEndpointApiKey, ok := os.LookupEnv("PolicyEndpointApiKey")

	if !ok {
		return Config{}, errors.New("not defined PolicyEndpointApiKey environment variable")
	}

	return Config{
		Url: policyEndpointUrl,
		Key: policyEndpointApiKey,
	}, nil

}

func NewClient() (Client, error) {

	config, configError := newConfigFromEnvironment()

	if configError != nil {
		return Client{}, errors.Wrap(configError, "could not read configuration from environment")
	}

	return Client{
		restClient: resty.New(),
		config:     config,
	}, nil

}

func (c Client) searchPolicies() (*http.Response, error) {

	searchUrl := c.config.Url

	searchResponse, error := c.restClient.
		SetHeader("x-api-key", c.config.Key).
		GetClient().
		Get(searchUrl)

	if error != nil {
		return nil, errors.Wrapf(error, "cannot search policies by rest client")
	}

	if searchResponse.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("error with status  %s", searchResponse.Status))
	}

	return searchResponse, nil

}
