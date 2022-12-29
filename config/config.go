package config

import (
	"errors"
	"feed-processor/integrators"
	"feed-processor/tenant"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const (
	configPath = "config.yaml"
)

// Config represents the structure of the YAML config file.
type Config struct {
	Tenants []*Tenant `yaml:"tenants"`
}

type Tenant struct {
	TenantID   string    `yaml:"tenant_id"`
	TenantName string    `yaml:"tenant_name"`
	Sources    []*Source `yaml:"sources"`
}

type Source struct {
	Type      string     `yaml:"type"`
	Discourse *Discourse `yaml:"discourse"`
	Twitter   *Twitter   `yaml:"twitter"`
	Intercom  *Intercom  `yaml:"intercom"`
	Playstore *Playstore `yaml:"playstore"`
}

type Discourse struct {
	BaseURL string `yaml:"base_url"`
}

type Twitter struct {
	ConsumerKey    string `yaml:"consumer_key"`
	ConsumerSecret string `yaml:"consumer_secret"`
	AccessToken    string `yaml:"access_token"`
	AccessSecret   string `yaml:"access_secret"`
}

type Intercom struct {
	AppID   string `yaml:"app_id"`
	APIKey  string `yaml:"api_key"`
	BaseURL string `yaml:"base_url"`
}

type Playstore struct {
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

func getSource(source Source) (integrators.Integrator, error) {
	switch source.Type {
	case "discourse":
		return integrators.NewDiscourseIntegrator(source.Discourse.BaseURL), nil
	case "twitter":
		return integrators.NewTwitterIntegrator(source.Twitter.ConsumerKey,
			source.Twitter.ConsumerSecret, source.Twitter.AccessToken,
			source.Twitter.AccessSecret), nil
	case "intercom":
		return integrators.NewIntercomIntegrator(source.Intercom.AppID,
			source.Intercom.APIKey, source.Intercom.BaseURL), nil
	case "playstore":
		return integrators.NewPlaystoreIntegrator(source.Playstore.AppPackageNames), nil
	default:
		return nil, errors.New(fmt.Sprintf("not valid source type: %s", source.Type))
	}
}
