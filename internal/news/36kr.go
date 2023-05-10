package news

import (
	"AliGenieServer/internal/model"
	"AliGenieServer/internal/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// https://gateway.36kr.com/api/mis/me/article

type KrReq struct {
	PartnerId string   `json:"partner_id"`
	Timestamp int64    `json:"timestamp"`
	Param     *KrParam `json:"param"`
}

type KrParam struct {
	UserId       string `json:"userId"`
	PageEvent    int    `json:"pageEvent"`
	PageSize     int    `json:"pageSize"`
	PageCallback string `json:"pageCallback"`
	SiteId       int    `json:"siteId"`
	PlatformId   int    `json:"platformId"`
}

type KrRes struct {
	Code int     `json:"code"`
	Data *KrData `json:"data"`
}

type KrData struct {
	PageCallback string    `json:"pageCallback"`
	HasNextPage  int       `json:"hasNextPage"`
	ItemList     []*KrItem `json:"itemList"`
}

type KrItem struct {
	ItemId           int64               `json:"itemId"`
	ItemType         int                 `json:"itemType"`
	TemplateMaterial *KrTemplateMaterial `json:"templateMaterial"`
	Route            string              `json:"route"`
}

type KrTemplateMaterial struct {
	ItemId        int64  `json:"itemId"`
	TemplateType  int    `json:"templateType"`
	WidgetImage   string `json:"widgetImage"`
	PublishTime   int64  `json:"publishTime"`
	WidgetTitle   string `json:"widgetTitle"`
	WidgetContent string `json:"widgetContent"`
}

func GetNews(c *gin.Context) {
	if err, res := get815(); err != nil {
		c.JSON(http.StatusOK, &model.SkillRes{
			ReturnCode: "0",
			ReturnValue: &model.ReturnValue{
				Reply:              err.Error(),
				ResultType:         "RESULT",
				ExecuteCode:        "SUCCESS",
				SkillDialogSession: &model.SkillDialogSession{SkillEndNluSession: true},
			},
		})
	} else {
		// 8点1氪丨
		msg := res.Data.ItemList[0].TemplateMaterial.WidgetTitle
		msg = strings.Replace(msg, "8点1氪丨", "以下内容来自\"8点1氪\"", -1)
		c.JSON(http.StatusOK, &model.SkillRes{
			ReturnCode: "0",
			ReturnValue: &model.ReturnValue{
				Reply:              msg,
				ResultType:         "RESULT",
				ExecuteCode:        "SUCCESS",
				SkillDialogSession: &model.SkillDialogSession{SkillEndNluSession: true},
			},
		})
	}
}

func get815() (error, *KrRes) {
	req := &KrReq{
		PartnerId: "web",
		Param: &KrParam{
			UserId:     "5652071",
			SiteId:     1,
			PlatformId: 2,
			PageSize:   1,
		},
	}
	body := utils.SendPost("https://gateway.36kr.com/api/mis/me/article", map[string]string{"Content-Type": "application/json"}, req)
	res := &KrRes{}
	err := json.Unmarshal(body, &res)
	return err, res
}
