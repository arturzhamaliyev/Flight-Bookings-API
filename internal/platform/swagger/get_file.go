package swagger

import (
	"os"

	"go.uber.org/zap"
)

// GetSwaggerYaml reads specification from yaml file and returns it as bytes.
func GetSwaggerYaml() []byte {
	f, err := os.ReadFile("./docs/docs.yaml")
	if err != nil {
		zap.S().Info(err)
		return nil
	}
	return f
}
