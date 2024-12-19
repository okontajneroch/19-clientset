package v1

import (
	starwarsv1 "github.com/okontajneroch/starwars/api/starwars/v1"
	"github.com/okontajneroch/starwars/clientset/scheme"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/rest"
)

// vytvorim si interface pre moju skupinu 'starwars.okontajneroch.sk/v1'
type StarwarsV1Interface interface {
	Starfighters(namespace string) StarfighterInterface
}

// vytvorim si strukturu, ktora bude implementovat interface skupiny.
// Tato struktura bude opat obsahovat referenciu na REST klienta
type StarwarsV1Client struct {
	restClient rest.Interface
}

// metoda, ktora vrati interface pre pracu so Starfighter objektami
func (c *StarwarsV1Client) Starfighters(namespace string) StarfighterInterface {
	return newStarfighters(c.restClient, namespace)
}

// factory metoda, ktora vytvori REST klienta pre skupinu 'starwars.okontajneroch.sk/v1'
// a vrati ho ako interface.
func NewForConfig(c *rest.Config) (*StarwarsV1Client, error) {
	config := *c // shallow copy
	config.APIPath = "/apis"
	config.GroupVersion = &starwarsv1.SchemeGroupVersion
	config.NegotiatedSerializer = serializer.NewCodecFactory(scheme.Scheme)

	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	return New(client), nil
}

func New(c rest.Interface) *StarwarsV1Client {
	return &StarwarsV1Client{restClient: c}
}
