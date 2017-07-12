package openshift

import "testing"

func TestProcessEvent(t *testing.T) {
	err := processEvents()
	if err != nil {
		t.Fatal(err)
	}
}
