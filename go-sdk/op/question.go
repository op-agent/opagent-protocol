package op

type QuestionOption struct {
	Label       string `json:"label"`
	Description string `json:"description,omitempty"`
}

type QuestionInfo struct {
	Question string           `json:"question"`
	Header   string           `json:"header"`
	Options  []QuestionOption `json:"options,omitempty"`
	Multiple bool             `json:"multiple,omitempty"`
	Custom   bool             `json:"custom,omitempty"`
}

type ElicitationQuestionRequest struct {
	RequestID    string         `json:"requestID"`
	Questions    []QuestionInfo `json:"questions"`
	CurrentIndex int            `json:"currentIndex,omitempty"`
}
