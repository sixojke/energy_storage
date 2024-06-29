package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) InitSolarPanelRoutes(api *gin.RouterGroup) {
	solarPanel := api.Group("/solar_panel")
	{
		solarPanel.GET("", h.solarPanelAllData)
	}
}

func (h *Handler) solarPanelAllData(c *gin.Context) {
	data, err := h.services.SolarPanel.AllData()
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, dataResponse{
		Data:  data,
		Count: len(data),
	})
}
