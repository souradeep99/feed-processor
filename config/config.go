package config

import (
	"errors"
	"feed-processor/integrators"
	"feed-processor/integrators/source/discourse"
	"feed-processor/integrators/source/intercom"
	"feed-processor/integrators/source/playstore"
	"feed-processor/integrators/source/twitter"
	"feed-processor/tenant"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const (
	// configPath is the path to the YAML config file.
	configPath = "config.yaml"
)

// Config represents the structure of the YAML config file.
type Config struct {
	Tenants []*Tenant `yaml:"tenants"`
}

type Tenant struct {
	// TenantID is the unique identifier for the tenant.
	TenantID string `yaml:"tenant_id"`
	// TenantName is the name of the tenant.
	TenantName string `yaml:"tenant_name"`
	// Sources is a list of feedback sources for the tenant.
	Sources []*Source `yaml:"sources"`
}

type Source struct {
	// Type is the type of feedback source, e.g. "discourse", "twitter", etc.
	Type string `yaml:"type"`
	// Discourse is the configuration for the Discourse feedback source.
	Discourse *Discourse `yaml:"discourse"`
	// Twitter is the configuration for the Twitter feedback source.
	Twitter *Twitter `yaml:"twitter"`
	// Intercom is the configuration for the Intercom feedback source.
	Intercom *Intercom `yaml:"intercom"`
	// Playstore is the configuration for the Playstore feedback source.
	Playstore *Playstore `yaml:"playstore"`
}

type Discourse struct {
	// BaseURL is the base URL of the Discourse forum.
	BaseURL string `yaml:"base_url"`
}

type Twitter struct {
	// ConsumerKey is the Twitter API consumer key.
	ConsumerKey string `yaml:"consumer_key"`
	// ConsumerSecret is the Twitter API consumer secret.
	ConsumerSecret string `yaml:"consumer_secret"`
	// AccessToken is the Twitter API access token.
	AccessToken string `yaml:"access_token"`
	// AccessSecret is the Twitter API access secret.
	AccessSecret string `yaml:"access_secret"`
}

type Intercom struct {
	// AppID is the Intercom app ID.
	AppID string `yaml:"app_id"`
	// APIKey is the Intercom API key.
	APIKey string `yaml:"api_key"`
	// BaseURL is the base URL of the Intercom API.
	BaseURL string `yaml:"base_url"`
}

type Playstore struct {
	// AppPackageNames is a list of package names for Playstore
	// apps that this tenant is interested in receiving feedback.
	AppPackageNames []string `yaml:"app_package_names"`
}

// LoadConfig loads the config from the YAML file.
func LoadConfig() (*Config, error) {
	// Read the YAML file.
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	// Unmarshal the YAML data into the Config struct.
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

// GetTenants returns a list of Tenant objects based on
// the configuration in the YAML config file.
func GetTenants() ([]*tenant.Tenant, error) {
	var tenants []*tenant.Tenant
	config, err := LoadConfig()
	if err != nil {
		return nil, err
	}
	if config == nil {
		return nil, errors.New("empty config")
	}
	for _, t := range config.Tenants {
		tenantData := &tenant.Tenant{
			ID:          t.TenantID,
			Name:        t.TenantName,
			Integrators: []integrators.Integrator{},
		}
		// Iterate over the sources for this tenant and
		// create the corresponding Integrator objects.
		for _, source := range t.Sources {
			if source == nil {
				return nil, errors.New("empty source")
			}
			integrator, err := getSource(*source)
			if err != nil {
				return nil, err
			}
			tenantData.Integrators = append(tenantData.Integrators, integrator)
		}
		tenants = append(tenants, tenantData)
	}
	return tenants, nil
}

// getSource creates an Integrator instance based on the given Source.
func getSource(source Source) (integrators.Integrator, error) {
	switch source.Type {
	case "discourse":
		if source.Discourse == nil {
			return nil, errors.New("empty discourse source")
		}
		return discourse.NewDiscourseIntegrator(source.Discourse.BaseURL), nil
	case "twitter":
		if source.Twitter == nil {
			return nil, errors.New("empty twitter source")
		}
		return twitter.NewTwitterIntegrator(source.Twitter.ConsumerKey,
			source.Twitter.ConsumerSecret, source.Twitter.AccessToken,
			source.Twitter.AccessSecret), nil
	case "intercom":
		if source.Intercom == nil {
			return nil, errors.New("empty intercom source")
		}
		return intercom.NewIntercomIntegrator(source.Intercom.AppID,
			source.Intercom.APIKey, source.Intercom.BaseURL), nil
	case "playstore":
		if source.Playstore == nil {
			return nil, errors.New("empty playstore source")
		}
		return playstore.NewPlaystoreIntegrator(source.Playstore.AppPackageNames), nil
	default:
		return nil, errors.New(fmt.Sprintf("not valid source type: %s", source.Type))
	}
}
