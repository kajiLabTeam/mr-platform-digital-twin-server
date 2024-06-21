package util

import (
	"encoding/json"
	"net/http"

	"github.com/kajiLabTeam/mr-platform-digital-twin-server/common"
	"github.com/kajiLabTeam/mr-platform-digital-twin-server/model"
)

func UpdateHtml2dContent(content common.Content) (common.ResponseHtml2d, error) {
	// contentIdを元にコンテンツを更新
	html2d := common.Html2d{}
	var err error
	// コンテンツが既にデータベースに存在しているか
	exists, err := model.IsExistHtml2dContent(content.ContentId)
	if err != nil {
		return common.ResponseHtml2d{}, err
	}

	if exists {
		html2d, err = model.GetHtml2dContent(content.ContentId)
		if err != nil {
			return common.ResponseHtml2d{}, err
		}
	}

	if !exists {
		// ない場合は新規作成
		requestHtml2d := common.RequestHtml2d{}
		requestHtml2d, err = getHtml2dContentFromServer(content.Domain)
		if err != nil {
			return common.ResponseHtml2d{}, err
		}
		html2d = common.Html2d{
			ContentId: content.ContentId,
			Location:  requestHtml2d.Location,
			Rotation:  requestHtml2d.Rotation,
			TextType:  requestHtml2d.TextType,
			TextURL:   requestHtml2d.TextURL,
			StyleURL:  requestHtml2d.StyleURL,
		}
		// DBに新規作成
		if err := model.CreateHtml2dContent(html2d); err != nil {
			return common.ResponseHtml2d{}, err
		}
	}

	responseHtml2d := common.ResponseHtml2d{
		ContentId: html2d.ContentId,
		Location:  html2d.Location,
		Rotation:  html2d.Rotation,
		TextType:  html2d.TextType,
		TextURL:   html2d.TextURL,
		StyleURL:  html2d.StyleURL,
	}
	return responseHtml2d, nil
}

func getHtml2dContentFromServer(domain string) (common.RequestHtml2d, error) {
	// Domain からコンテンツを取得
	endpoint := domain + "/api/get/content/html2d"
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return common.RequestHtml2d{}, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return common.RequestHtml2d{}, err
	}
	defer resp.Body.Close()

	html2d := common.RequestHtml2d{}
	err = json.NewDecoder(resp.Body).Decode(&html2d)
	if err != nil {
		return common.RequestHtml2d{}, err
	}
	return html2d, nil
}
