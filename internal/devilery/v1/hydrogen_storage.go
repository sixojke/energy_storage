package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) InitHydrogenStorageRoutes(api *gin.RouterGroup) {
	hydrogenStorage := api.Group("/hydrogen_storage")
	{
		hydrogenStorage.GET("", h.hydrogenStorageAllData)
	}
}

func (h *Handler) hydrogenStorageAllData(c *gin.Context) {
	data, err := h.services.HydrogenStorage.AllData()
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, dataResponse{
		Data:  data,
		Count: len(data),
	})
}
