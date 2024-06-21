package util

import (
	"encoding/json"
	"net/http"

	"github.com/kajiLabTeam/mr-platform-digital-twin-server/common"
	"github.com/kajiLabTeam/mr-platform-digital-twin-server/model"
)

func UpdateModel3dContent(content common.Content) (common.ResponseModel3d, error) {
	// contentIdを元にコンテンツを更新
	model3d := common.Model3d{}
	var err error
	exists, err := model.IsExistModel3dContent(content.ContentId)
	if err != nil {
		return common.ResponseModel3d{}, err
	}

	if exists {
		model3d, err = model.GetModel3dContent(content.ContentId)
		if err != nil {
			return common.ResponseModel3d{}, err
		}
	}

	if !exists {
		// ない場合は新規作成
		requestModel3d := common.RequestModel3d{}
		requestModel3d, err = getModel3dContentFromServer(content.Domain)
		if err != nil {
			return common.ResponseModel3d{}, err
		}
		model3d = common.Model3d{
			ContentId: content.ContentId,
			Location:  requestModel3d.Location,
			Rotation:  requestModel3d.Rotation,
		}
		// DBに新規作成
		if err := model.CreateModel3dContent(model3d); err != nil {
			return common.ResponseModel3d{}, err
		}
	}

	// minioの署名付きURLの発行
	presignedURL, err := GetContentUrl(model3d.PresignedURL)
	if err != nil {
		return common.ResponseModel3d{}, err
	}
	model3d.PresignedURL = presignedURL

	responseModel3d := common.ResponseModel3d{
		ContentId:    model3d.ContentId,
		Location:     model3d.Location,
		Rotation:     model3d.Rotation,
		PresignedURL: model3d.PresignedURL,
	}
	return responseModel3d, nil
}

func getModel3dContentFromServer(domain string) (common.RequestModel3d, error) {
	// Domain からコンテンツを取得
	endpoint := domain + "/api/get/content/model3d"
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return common.RequestModel3d{}, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return common.RequestModel3d{}, err
	}
	defer resp.Body.Close()
	model3d := common.RequestModel3d{}
	err = json.NewDecoder(resp.Body).Decode(&model3d)
	if err != nil {
		return common.RequestModel3d{}, err
	}
	return model3d, nil
}
