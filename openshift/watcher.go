package openshift

import (
	"context"
	"encoding/json"
	"github.com/astaxie/beego"
	"k8s.io/client-go/rest"
	"net/http"
	"sync"
	"time"
)

type WatchFunc func(r Resource, obj interface{})

type Watcher struct {
	restClient *rest.RESTClient
	resource   Resource
	host       string
	cancel     context.CancelFunc
	running    bool
	mutex      sync.Mutex
	handler    WatchFunc
}

func NewWatcher(host string, rc *rest.RESTClient, r Resource) *Watcher {
	w := &Watcher{host: host, restClient: rc, resource: r}
	return w
}

func (w *Watcher) Start(handler WatchFunc) {
	if w.running {
		return
	}
	w.running = true
	w.handler = handler
	go func() {
		beego.Info("Watcher for", w.host, w.resource.Name, "started")
		defer beego.Info("Watch for", w.host, w.resource.Name, "ended")
		for w.running {
			err := w.processEvents()
			if err != nil {
				beego.Error(err)
				time.Sleep(time.Second * 10)
			}
		}
	}()
}

func (w *Watcher) Stop() {
	beego.Info("Stopping watcher", w.host, w.resource)
	w.running = false
	if w.cancel != nil {
		w.cancel()
	}

}

func (w *Watcher) processEvents() error {
	beego.Debug("Watching resource ", w.resource)
	var cx context.Context
	cx, w.cancel = context.WithCancel(context.Background())
	req, err := http.NewRequest(http.MethodGet, w.host+w.resource.WatchURL, nil)
	if err != nil {
		return err
	}
	req = req.WithContext(cx)
	resp, err := w.restClient.Client.Do(req)
	if err != nil {
		beego.Error("Request", req, "failed with", err)
		return err
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	for {
		var obj map[string]interface{}
		err = decoder.Decode(&obj)
		if err != nil {
			beego.Error("Failed to decode object from response:", err)
			return err
		}
		go w.handleEvent(obj)
	}
	return err
}

func (w *Watcher) handleEvent(obj interface{}) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	beego.Debug("Handing object", obj)
	if w.handler != nil {
		go w.handler(w.resource, obj)
	}
}
