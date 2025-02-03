package featuresyaml

import (
	"io"
	"os"

	"github.com/ghodss/yaml"
	"github.com/loft-sh/admin-apis/pkg/licenseapi"
)

func ReadFeaturesYaml(yamlPath string) ([]*licenseapi.Feature, error) {
	featuresYaml, err := os.Open(yamlPath)
	if err != nil {
		panic(err)
	}

	bytes, err := io.ReadAll(featuresYaml)
	if err != nil {
		panic(err)
	}

	return featuresUnmarshal(bytes)
}

func featuresUnmarshal(body []byte) ([]*licenseapi.Feature, error) {
	features := struct {
		Features []*licenseapi.Feature `json:"features"`
	}{}
	err := yaml.Unmarshal(body, &features)
	if err != nil {
		return nil, err
	}
	return features.Features, nil
}
