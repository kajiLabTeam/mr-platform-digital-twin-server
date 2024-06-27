package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/mr-platform-digital-twin-server/common"
	"github.com/kajiLabTeam/mr-platform-digital-twin-server/model"
	"github.com/kajiLabTeam/mr-platform-digital-twin-server/util"
)

func UpdateSpace(c *gin.Context) {
	// jsonをパース 型は common.RelayServerRequest
	var req common.RequestRelayServer
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 配列の中身が空の場合は204を返す
	if len(req.ContentIds) == 0 {
		c.Status(http.StatusNoContent)
		return
	}

	// 生成された content を格納するための配列
	var responseContents common.ResponseContents

	// 既存のコンテンツにないcontentIdの場合は新規作成
	for _, contentId := range req.ContentIds {
		content, err := model.GetContent(contentId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		err = procContent(content, content.ContentType, &responseContents)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	// responseContents.ResponseHtml2ds と responseContents.ResponseModel3ds が空の場合は空の配列を返す
	if responseContents.ResponseHtml2ds == nil {
		responseContents.ResponseHtml2ds = []common.ResponseHtml2d{}
	}
	if responseContents.ResponseModel3ds == nil {
		responseContents.ResponseModel3ds = []common.ResponseModel3d{}
	}

	c.JSON(http.StatusOK, responseContents)
}

func procContent(content common.Content, contentType string, responseContents *common.ResponseContents) error {
	if contentType == "html2d" {
		err := procHtml2dContent(content, responseContents)
		if err != nil {
			return err
		}
	} else if contentType == "model3d" {
		err := procModel3dContent(content, responseContents)
		if err != nil {
			return err
		}
	}
	return nil
}

func procHtml2dContent(content common.Content, responseContents *common.ResponseContents) error {
	responseHtml2d, err := util.UpdateHtml2dContent(content)
	if err != nil {
		return err
	}
	// Convert html2dContent to ResponseHtml2d before appending
	responseContents.ResponseHtml2ds = append(responseContents.ResponseHtml2ds, responseHtml2d)
	return nil
}

func procModel3dContent(content common.Content, responseContents *common.ResponseContents) error {
	responseModel3d, err := util.UpdateModel3dContent(content)
	if err != nil {
		return err
	}
	responseContents.ResponseModel3ds = append(responseContents.ResponseModel3ds, responseModel3d)
	return nil
}
