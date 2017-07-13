package openshift

import "testing"

func TestLoadResourcesFromSwagger(t *testing.T) {
	err := LoadResourcesFromSwagger("/home/danny/src/go/src/github.com/openshift/origin/api/swagger-spec/openshift-openapi-spec.json")
	if err != nil {
		t.Fatal(err)
	}
}
