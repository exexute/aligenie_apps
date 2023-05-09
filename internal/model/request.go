package model

import "encoding/json"

type SkillReq struct {
	SessionId   string `json:"sessionId"`
	Utterance   string `json:"utterance"`
	Token       string `json:"token"`
	RequestData struct {
		UserOpenId     string `json:"userOpenId"`
		DeviceOpenId   string `json:"deviceOpenId"`
		City           string `json:"city"`
		DeviceUnionIds string `json:"deviceUnionIds"`
		UserUnionIds   string `json:"userUnionIds"`
	} `json:"requestData"`
	BotId        int    `json:"botId"`
	DomainId     int    `json:"domainId"`
	SkillId      int    `json:"skillId"`
	SkillName    string `json:"skillName"`
	IntentId     int    `json:"intentId"`
	IntentName   string `json:"intentName"`
	SlotEntities []struct {
		IntentParameterId   int    `json:"intentParameterId"`
		IntentParameterName string `json:"intentParameterName"`
		OriginalValue       string `json:"originalValue"`
		StandardValue       string `json:"standardValue"`
		LiveTime            int    `json:"liveTime"`
		CreateTimeStamp     int64  `json:"createTimeStamp"`
		SlotName            string `json:"slotName"`
		SlotValue           string `json:"slotValue"`
	} `json:"slotEntities"`
	RequestId string `json:"requestId"`
	Device    struct {
		DeviceOpenId string `json:"deviceOpenId"`
	} `json:"device"`
	SkillSession struct {
		SkillSessionId string `json:"skillSessionId"`
		NewSession     bool   `json:"newSession"`
	} `json:"skillSession"`
}

func (receiver *SkillReq) String() string {
	data, _ := json.Marshal(receiver)
	return string(data)
}
