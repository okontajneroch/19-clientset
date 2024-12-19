package main

import (
	"context"
	"fmt"

	starwarsv1 "github.com/okontajneroch/starwars/api/starwars/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	clientsetv1 "github.com/okontajneroch/starwars/generated"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, _ := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		nil,
	).ClientConfig()

	client, _ := clientsetv1.NewForConfig(config)

	xwing := &starwarsv1.Starfighter{}
	xwing.Name = "x-wing-1"
	xwing.Spec.Pilot = "Luke Skywalker"
	xwing.Spec.Type = "X-Wing"
	xwing.Spec.Faction = starwarsv1.Rebellion

	result, err := client.
		StarwarsV1().
		Starfighters("default").
		Create(context.Background(), xwing, metav1.CreateOptions{})

	if err != nil {
		fmt.Printf("Error creating Starfighter: %s\n", err)
		return
	}

	fmt.Printf("Created Starfighter: %s %s \n", result.Name, result.UID)
}
