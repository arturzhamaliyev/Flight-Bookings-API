package swagger

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/platform/config"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/platform/swagger"
	"github.com/flowchartsman/swaggerui"
)

// New hosts swagger documentation on separate port.
func New(cfg config.Config) {
	zap.S().Infof("swagger server listening on port: %v", cfg.Swagger.Port)
	err := http.ListenAndServe(":"+cfg.Swagger.Port, swaggerui.Handler(swagger.GetSwaggerYaml()))
	if err != nil {
		zap.S().Info(err)
	}
}
