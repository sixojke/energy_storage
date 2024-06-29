package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) InitThermalStorageRoutes(api *gin.RouterGroup) {
	thermalStorage := api.Group("/thermal_storage")
	{
		thermalStorage.GET("", h.thermalStorageAllData)
	}
}

func (h *Handler) thermalStorageAllData(c *gin.Context) {
	data, err := h.services.ThermalStorage.AllData()
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, dataResponse{
		Data:  data,
		Count: len(data),
	})
}
