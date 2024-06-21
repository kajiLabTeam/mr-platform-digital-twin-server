package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/mr-platform-digital-twin-server/common"
	"github.com/kajiLabTeam/mr-platform-digital-twin-server/model"
)

func CreatePublicSpaceContent(c *gin.Context) {
	// jsonをパース 型は common.Content
	var requestContent common.RequestContent
	if err := c.ShouldBindJSON(&requestContent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 既存のコンテンツにないcontentIdの場合は新規作成
	exists, err := model.IsExistContent(requestContent.ContentType, requestContent.Domain)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if exists {
		contentId, err := model.GetContentId(requestContent.ContentType, requestContent.Domain)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"contentId": contentId})
		return
	}

	// handle case where content does not exist
	contentId, err := model.CreateContent(requestContent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"contentId": contentId})
}
