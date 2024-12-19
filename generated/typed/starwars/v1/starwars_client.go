// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	http "net/http"

	starwarsv1 "github.com/okontajneroch/starwars/api/starwars/v1"
	scheme "github.com/okontajneroch/starwars/generated/scheme"
	rest "k8s.io/client-go/rest"
)

type StarwarsV1Interface interface {
	RESTClient() rest.Interface
	StarfightersGetter
}

// StarwarsV1Client is used to interact with features provided by the starwars.okontajneroch.sk group.
type StarwarsV1Client struct {
	restClient rest.Interface
}

func (c *StarwarsV1Client) Starfighters(namespace string) StarfighterInterface {
	return newStarfighters(c, namespace)
}

// NewForConfig creates a new StarwarsV1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*StarwarsV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new StarwarsV1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*StarwarsV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &StarwarsV1Client{client}, nil
}

// NewForConfigOrDie creates a new StarwarsV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *StarwarsV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new StarwarsV1Client for the given RESTClient.
func New(c rest.Interface) *StarwarsV1Client {
	return &StarwarsV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := starwarsv1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = rest.CodecFactoryForGeneratedClient(scheme.Scheme, scheme.Codecs).WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *StarwarsV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
