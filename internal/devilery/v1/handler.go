package v1

import (
	"energy_storage/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("v1")
	{
		h.initElectrochemicalBatteryRoutes(v1)
		h.InitThermalStorageRoutes(v1)
		h.InitHydrogenStorageRoutes(v1)
		h.InitWindTurbineRoutes(v1)
		h.InitSolarPanelRoutes(v1)
	}
}
