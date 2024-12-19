package clientset

import (
	starwarsv1api "github.com/okontajneroch/starwars/api/starwars/v1"
	"github.com/okontajneroch/starwars/clientset/scheme"
	starwarsv1 "github.com/okontajneroch/starwars/clientset/typed/starwars/v1"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/rest"
)

// Vytvorim si interface, ktory bude poskytovat metody na pristup ku clientom
// vsetkych skupin a verzii
type Interface interface {
	StarwarsV1() starwarsv1.StarwarsV1Interface
}

// Vytvorim si strukturu Clientset, ktora je implementaciou Interface,
// a ktora bude zdruzovat vsetkych clientov
type Clientset struct {
	starwarsV1 *starwarsv1.StarwarsV1Client
}

func (c *Clientset) StarwarsV1() starwarsv1.StarwarsV1Interface {
	return c.starwarsV1
}

func NewForConfig(c *rest.Config) (*Clientset, error) {
	config := *c // shallow copy
	config.APIPath = "/apis"
	config.GroupVersion = &starwarsv1api.SchemeGroupVersion
	config.NegotiatedSerializer = serializer.NewCodecFactory(scheme.Scheme)

	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	cs := &Clientset{
		starwarsV1: starwarsv1.New(client),
	}

	return cs, nil
}
