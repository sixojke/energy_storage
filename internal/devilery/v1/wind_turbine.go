package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) InitWindTurbineRoutes(api *gin.RouterGroup) {
	windTurbine := api.Group("/wind_turbine")
	{
		windTurbine.GET("", h.windTurbineAllData)
	}
}

func (h *Handler) windTurbineAllData(c *gin.Context) {
	data, err := h.services.WindTurbine.AllData()
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, dataResponse{
		Data:  data,
		Count: len(data),
	})
}
