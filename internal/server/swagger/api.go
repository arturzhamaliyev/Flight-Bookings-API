package swagger

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/config"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/platform/swagger"
	"github.com/flowchartsman/swaggerui"
)

// New hosts swagger documentation on separate port.
func New(cfg config.Config) {
	zap.S().Infof("swagger server listening on port: %v", cfg.SwaggerPort)
	err := http.ListenAndServe(":"+cfg.SwaggerPort, swaggerui.Handler(swagger.GetSwaggerYaml()))
	if err != nil {
		zap.S().Info(err)
	}
}
