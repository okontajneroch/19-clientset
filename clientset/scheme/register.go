package scheme

import (
	starwarsv1 "github.com/okontajneroch/starwars/api/starwars/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// Túto schému budem používať v mojich clientsetoch na serializáciu
// a deserializáciu objektov.
var Scheme = runtime.NewScheme()

// ParameterCodec bude používaný na serializáciu a deserializáciu
// parametrov v URL
var ParameterCodec = runtime.NewParameterCodec(Scheme)

func init() {
	// Do schémy pridám všetky objekty, ktoré clientsety budu poznať
	// V tomto prípade celú skupinu starwars.okontajneroch.sk/v1
	starwarsv1.AddToScheme(Scheme)
}
