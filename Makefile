.PHONY: generate
generate:
	deepcopy-gen ./api/starwars/v1 --output-file zz_generated.deepcopy.go
	controller-gen crd paths=./api/... output:crd:dir=./k8s
