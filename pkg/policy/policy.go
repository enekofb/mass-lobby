package policy

import (
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"os"
)

type Client struct {
	restClient *resty.Client
	config     Config
}

type Config struct {
	Url string
}

func newConfigFromEnvironment() (Config, error) {

	policyEndpointUrl, ok := os.LookupEnv("PolicyEndpointUrl")

	if !ok {
		return Config{
			Url: policyEndpointUrl,
		}, errors.New("not defined PolicyEndpointUrl environment variable")

	}

	return Config{
		Url: policyEndpointUrl,
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

func (c Client) searchPolicies() {

	searchUrl := c.config.Url

	searchResponse, error := c.restClient.GetClient().Get(searchUrl)

	if error != nil {
		return Client{}, errors.Wrap(configError, "could not read configuration from environment")
	}

}
