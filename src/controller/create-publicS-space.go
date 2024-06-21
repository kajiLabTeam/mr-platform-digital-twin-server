package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/mr-platform-digital-twin-server/common"
	"github.com/kajiLabTeam/mr-platform-digital-twin-server/model"
)

func CreatePublicSpace(c *gin.Context) {
	var requestContent common.RequestPublicSpace
	if err := c.ShouldBindJSON(&requestContent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exists, err := model.IsExistPublicSpaces(requestContent.OrganizationId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if exists {
		publicSpaceId, err := model.GetPublicSpaceId(requestContent.OrganizationId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"publicSpaceId": publicSpaceId})
		return
	}

	publicSpaceId, err := model.CreatePublicSpaces(requestContent.OrganizationId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"publicSpaceId": publicSpaceId})
}
