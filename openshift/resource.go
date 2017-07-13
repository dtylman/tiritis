package openshift

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/swagger"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type Resource struct {
	Name     string
	WatchURL string
}

var Resources []Resource

func ResourceByName(name string) *Resource {
	for i := range Resources {
		if Resources[i].Name == name {
			return &Resources[i]
		}
	}
	return nil
}

//LoadResourcesFromSwagger loads resource from a openshift official api swagger file
func LoadResourcesFromSwagger(swaggerFile string) error {
	Resources = make([]Resource, 0)
	data, err := ioutil.ReadFile(swaggerFile)
	if err != nil {
		return err
	}
	var swag swagger.Swagger
	err = json.Unmarshal(data, &swag)
	if err != nil {
		return err
	}
	for path, item := range swag.Paths {
		if !strings.Contains(path, `/oapi/v1/`) && !strings.Contains(path, `/api/v1`) {
			continue
		}
		if strings.Contains(path, `{namespace}`) {
			continue
		}
		if item.Get != nil {
			for _, param := range item.Get.Parameters {
				if param.Name == "watch" {
					r := Resource{Name: filepath.Base(path), WatchURL: fmt.Sprintf("%s?watch=true", path)}
					Resources = append(Resources, r)
					beego.Debug("Loaded resource", r)
				}
			}

		}
	}
	return nil
}
