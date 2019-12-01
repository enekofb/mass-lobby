package aws

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
)

type Config struct {
	awsConfig aws.Config
}

func (c Config) Load() aws.Config {
	return c.awsConfig
}

func NewDefaultConfig() (Config, error) {

	defaultConfig, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return Config{}, err
	}
	return Config{awsConfig: defaultConfig}, nil
}
