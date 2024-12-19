package v1

import (
	"context"
	starwarsv1 "github.com/okontajneroch/starwars/api/starwars/v1"
	"github.com/okontajneroch/starwars/clientset/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

// Najprv si vytvorim interface clienta, so všetkými operáciami,
// ktoré chcem vykonávať nad objektami Starfighter
type StarfighterInterface interface {
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*starwarsv1.Starfighter, error)
	Create(ctx context.Context, starfighter *starwarsv1.Starfighter) (*starwarsv1.Starfighter, error)
	Update(ctx context.Context, starfighter *starwarsv1.Starfighter) (*starwarsv1.Starfighter, error)
}

// Potom si vytvorim štruktúru, ktorá bude implementovať interface
// StarfighterInterface. Táto štruktúra bude obsahovať restClienta
// a namespace, nad ktorym sa budú vykonávať operácie.
type starfighters struct {
	restClient rest.Interface
	ns         string
}

func newStarfighters(c rest.Interface, namespace string) *starfighters {
	return &starfighters{
		restClient: c,
		ns:         namespace,
	}
}

// nasledne si vytvorim implementacie metod pre interface
func (c *starfighters) Get(ctx context.Context, name string, opts metav1.GetOptions) (*starwarsv1.Starfighter, error) {
	result := &starwarsv1.Starfighter{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("starfighters").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return result, err
}

func (c *starfighters) Create(ctx context.Context, starfighter *starwarsv1.Starfighter) (*starwarsv1.Starfighter, error) {
	result := &starwarsv1.Starfighter{}
	err := c.restClient.
		Post().
		Namespace(c.ns).
		Resource("starfighters").
		Body(starfighter).
		Do(ctx).
		Into(result)
	return result, err
}

func (c *starfighters) Update(ctx context.Context, starfighter *starwarsv1.Starfighter) (*starwarsv1.Starfighter, error) {
	result := &starwarsv1.Starfighter{}
	err := c.restClient.
		Put().
		Namespace(c.ns).
		Resource("starfighters").
		Name(starfighter.Name).
		Body(starfighter).
		Do(ctx).
		Into(result)
	return result, err
}
