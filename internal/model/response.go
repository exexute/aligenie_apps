package model

type SkillRes struct {
	ReturnCode          string       `json:"returnCode"`
	ReturnErrorSolution string       `json:"returnErrorSolution"`
	ReturnMessage       string       `json:"returnMessage"`
	ReturnValue         *ReturnValue `json:"returnValue"`
}

type ReturnValue struct {
	Reply              string              `json:"reply"`       // 回复给用户的 TTS 文本信息
	ResultType         string              `json:"resultType"`  // 回复时的状态标识，RESULT: 天猫精灵播放完回复内容后不会开麦、ASK_INF: 引导继续对话/参数追问响应(播放完后开麦)、CONFIRM: 重要信息再次确认响应
	ExecuteCode        string              `json:"executeCode"` // SUCCESS: 代表执行成功、PARAMS_ERROR: 代表接收到的请求参数出错、EXECUTE_ERROR: 代表自身代码有异常、REPLY_ERROR: 代表回复结果生成出错
	Actions            []*Action           `json:"actions,omitempty"`
	SkillDialogSession *SkillDialogSession `json:"skillDialogSession,omitempty"` // 对话结束响应，技能的对话 session 控制
	AskedInfos         []*AskedInfo        `json:"askedInfos,omitempty"`         // 参数追问响应
	ConfirmParaInfo    *ConfirmParaInfo    `json:"confirmParaInfo,omitempty"`    // 配置 confirm 模型优先匹配内容的参数名称，用户的回答优先匹配此处定义的参数取值。
	SelectParaInfo     *SelectParaInfo     `json:"selectParaInfo,omitempty"`     // select状态下需要填充多个候选信息
	GwCommands         []*GwCommand        `json:"gwCommands,omitempty"`         // 在单论或连续对话中，有时我们要根据用户的语料来判定是否需要开麦，如 “天猫精灵 开始录音”然后天猫精灵就可以直接开麦了，不需要在回复用户，这时可以用 Listen 指令。
}

type SkillDialogSession struct {
	SkillEndNluSession bool              `json:"skillEndNluSession"` // 是否清除当前对话的 session，默认是 false
	SessionAttrs       map[string]string `json:"sessionAttrs"`       // 当前 session 中可传入下一轮会话的临时数据，会在下一轮对话的请求数据中携带。使用时必须设置 skillEndNluSession=false
}

type AskedInfo struct {
	ParameterName string `json:"parameterName"` // 追问的参数名称。此名称是在意图中定义的，不是实体标识
	IntentId      string `json:"intentId"`      // 参数所在的意图ID。线上领域和测试领域中的意图ID不同，要从请求数据中获取意图ID，请勿使用固定值
}

type ConfirmParaInfo struct {
	ConfirmParameterName string `json:"confirmParameterName"` // 用来匹配用户回答表示确定内容的参数名称
	DenyParameterName    string `json:"denyParameterName"`    // 用来匹配用户回答表示否定内容的参数名称
}

type Action struct {
	Name       string            `json:"name"`
	Properties map[string]string `json:"properties"`
}

type SelectParaInfo struct {
	IntentParameterName string             `json:"intentParameterName"` // 动态参数的参数名称
	ParameterReplyMatch bool               `json:"parameterReplyMatch"` // 是否匹配待选内容外的参数实体，默认false（不匹配）
	SelectIndexMatch    bool               `json:"selectIndexMatch"`    // 用户的选项 index 是否可以超出待选项个数范围，默认是false（不能超出）
	CandidateList       []*SelectCandidate `json:"candidateList"`       // 动态参数的待选项内容集合
}

type SelectCandidate struct {
	Value       string   `json:"value"`       // 每个待选项的候选值
	AliasValues []string `json:"aliasValues"` // 候选值的同义词列表
	NormValue   string   `json:"normValue"`   // 候选归一化值，默认为空
}

type GwCommand struct {
	CommandDomain string `json:"commandDomain"`
	CommandName   string `json:"commandName"`
	Payload       struct {
		NeedLight bool `json:"needLight"`
		NeedVoice bool `json:"needVoice"`
	} `json:"payload"`
}
