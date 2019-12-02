package policy

import "github.com/go-resty/resty/v2"

type PolicyClient struct {
	client *resty.Client
}

func NewPolicyClient() (PolicyClient, error) {

	return PolicyClient{
		client: resty.New(),
	}, nil

}
