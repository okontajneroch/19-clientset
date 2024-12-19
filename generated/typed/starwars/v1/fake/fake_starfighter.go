// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/okontajneroch/starwars/api/starwars/v1"
	starwarsv1 "github.com/okontajneroch/starwars/generated/typed/starwars/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeStarfighters implements StarfighterInterface
type fakeStarfighters struct {
	*gentype.FakeClientWithList[*v1.Starfighter, *v1.StarfighterList]
	Fake *FakeStarwarsV1
}

func newFakeStarfighters(fake *FakeStarwarsV1, namespace string) starwarsv1.StarfighterInterface {
	return &fakeStarfighters{
		gentype.NewFakeClientWithList[*v1.Starfighter, *v1.StarfighterList](
			fake.Fake,
			namespace,
			v1.SchemeGroupVersion.WithResource("starfighters"),
			v1.SchemeGroupVersion.WithKind("Starfighter"),
			func() *v1.Starfighter { return &v1.Starfighter{} },
			func() *v1.StarfighterList { return &v1.StarfighterList{} },
			func(dst, src *v1.StarfighterList) { dst.ListMeta = src.ListMeta },
			func(list *v1.StarfighterList) []*v1.Starfighter { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.StarfighterList, items []*v1.Starfighter) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}
