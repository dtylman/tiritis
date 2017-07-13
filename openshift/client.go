package openshift

import (
	"errors"
	"github.com/astaxie/beego"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"sync"
)

//DefaultClient is just default client for now.
var DefaultClient *Client

type Client struct {
	watchers   map[string]*Watcher
	restClient *rest.RESTClient
	config     *rest.Config
	mutex      sync.Mutex
}

func NewClient() (*Client, error) {
	c := new(Client)
	c.watchers = make(map[string]*Watcher)
	var err error
	c.config, err = getConfig()
	if err != nil {
		return nil, err
	}
	c.restClient, err = rest.UnversionedRESTClientFor(c.config)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Client) WatchResource(r Resource, handler WatchFunc) *Watcher {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	w, exists := c.watchers[r.Name]
	if !exists {
		w = NewWatcher(c.config.Host, c.restClient, r)
		w.Start(handler)
		c.watchers[r.Name] = w
	}
	return w
}

func (c *Client) Stop() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for _, w := range c.watchers {
		w.Stop()
	}
	c.watchers = make(map[string]*Watcher)
}

func getConfig() (*rest.Config, error) {
	var config *rest.Config
	var err error
	kubeConfig := os.Getenv("KUBE_CONFIG")
	if kubeConfig != "" {
		config, err = clientcmd.BuildConfigFromFlags("", kubeConfig)
		if err != nil {
			return nil, err
		}
		err = rest.SetKubernetesDefaults(config)
		if err != nil {
			return nil, err
		}
		config.NegotiatedSerializer = scheme.Codecs
	} else {
		config, err = rest.InClusterConfig()
		if err != nil {
			return nil, errors.New("KUBE_CONFIG is not defined AND for in-cluster config " + err.Error())
		}
		return config, err
	}
	beego.Info("Loaded kube config for:", config.Host)
	return config, nil
}
