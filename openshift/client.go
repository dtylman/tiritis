package openshift

import (
	"encoding/json"
	"fmt"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func processEvents() error {
	// creates the in-cluster config
	config, err := clientcmd.BuildConfigFromFlags("", "/home/danny/.kube/config")
	if err != nil {
		return err
	}
	err = rest.SetKubernetesDefaults(config)
	if err != nil {
		return err
	}
	config.NegotiatedSerializer = scheme.Codecs

	rc, err := rest.UnversionedRESTClientFor(config)
	if err != nil {
		return err
	}

	resp, err := rc.Client.Get(config.Host + "/oapi/v1/routes")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var obj map[string]interface{}
	for err == nil {
		err = decoder.Decode(&obj)
		output, _ := json.Marshal(obj)
		fmt.Println(string(output))
	}
	return err

}
