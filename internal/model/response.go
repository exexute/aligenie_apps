package model

type SkillRes struct {
	ReturnCode          string      `json:"returnCode"`
	ReturnErrorSolution string      `json:"returnErrorSolution"`
	ReturnMessage       string      `json:"returnMessage"`
	ReturnValue         interface{} `json:"returnValue"`
}

type EndSessionReturnValue struct {
	Reply              string              `json:"reply"`
	ResultType         string              `json:"resultType"`
	SkillDialogSession *SkillDialogSession `json:"skillDialogSession"`
	ExecuteCode        string              `json:"executeCode"`
}

type SkillDialogSession struct {
	SkillEndNluSession bool `json:"skillEndNluSession"`
}
