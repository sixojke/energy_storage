package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initElectrochemicalBatteryRoutes(api *gin.RouterGroup) {
	electrochemicalBattery := api.Group("/electrochemical_battery")
	{
		electrochemicalBattery.GET("", h.electrochemicalBatteryallData)
	}
}

func (h *Handler) electrochemicalBatteryallData(c *gin.Context) {
	data, err := h.services.ElectrochemicalBattery.AllData()
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, dataResponse{
		Data:  data,
		Count: len(data),
	})
}
