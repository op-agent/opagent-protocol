package op

type ChatReviewFileStatus string

const (
	ChatReviewFilePending    ChatReviewFileStatus = "pending"
	ChatReviewFileApproved   ChatReviewFileStatus = "approved"
	ChatReviewFileRejected   ChatReviewFileStatus = "rejected"
	ChatReviewFileRolledBack ChatReviewFileStatus = "rolledBack"
)

type ChatReviewTurnStatus string

const (
	ChatReviewTurnPending    ChatReviewTurnStatus = "pending"
	ChatReviewTurnResolved   ChatReviewTurnStatus = "resolved"
	ChatReviewTurnRolledBack ChatReviewTurnStatus = "rolledBack"
)

type ChatReviewDecision string

const (
	ChatReviewDecisionApprove    ChatReviewDecision = "approve"
	ChatReviewDecisionReject     ChatReviewDecision = "reject"
	ChatReviewDecisionApproveAll ChatReviewDecision = "approveAll"
	ChatReviewDecisionRejectAll  ChatReviewDecision = "rejectAll"
)

type ChatReviewRollbackScope string

const (
	ChatReviewRollbackFile ChatReviewRollbackScope = "file"
	ChatReviewRollbackTurn ChatReviewRollbackScope = "turn"
)

type ChatReviewLineRange struct {
	StartLine int `json:"startLine"`
	EndLine   int `json:"endLine"`
}

type ChatReviewFile struct {
	Path               string                `json:"path"`
	Status             ChatReviewFileStatus  `json:"status"`
	Diff               string                `json:"diff"`
	BaselineExists     bool                  `json:"baselineExists"`
	FirstChangedLine   int                   `json:"firstChangedLine,omitempty"`
	FirstChangedColumn int                   `json:"firstChangedColumn,omitempty"`
	LineCount          int                   `json:"lineCount,omitempty"`
	ChangedRanges      []ChatReviewLineRange `json:"changedRanges,omitempty"`
}

type ChatReviewState struct {
	ThreadID        string               `json:"threadID"`
	TurnID          string               `json:"turnID"`
	ChatPath        string               `json:"chatPath"`
	Status          ChatReviewTurnStatus `json:"status"`
	CreatedAt       string               `json:"createdAt"`
	CanReview       bool                 `json:"canReview"`
	CanRollback     bool                 `json:"canRollback"`
	Unresolved      int                  `json:"unresolved"`
	ApprovedCount   int                  `json:"approvedCount"`
	RejectedCount   int                  `json:"rejectedCount"`
	RolledBackCount int                  `json:"rolledBackCount"`
	Files           []ChatReviewFile     `json:"files"`
}

type ChatReviewListParams struct {
	ThreadID string `json:"threadID,omitempty"`
	ChatPath string `json:"chatPath,omitempty"`
}

type ChatReviewListResult struct {
	Reviews []ChatReviewState `json:"reviews,omitempty"`
}

type ChatReviewResolveParams struct {
	ThreadID string             `json:"threadID,omitempty"`
	ChatPath string             `json:"chatPath,omitempty"`
	TurnID   string             `json:"turnID"`
	Decision ChatReviewDecision `json:"decision"`
	Path     string             `json:"path,omitempty"`
}

type ChatReviewResolveResult struct {
	Review *ChatReviewState `json:"review,omitempty"`
}

type ChatReviewRollbackParams struct {
	ThreadID string                  `json:"threadID,omitempty"`
	ChatPath string                  `json:"chatPath,omitempty"`
	TurnID   string                  `json:"turnID"`
	Scope    ChatReviewRollbackScope `json:"scope"`
	Path     string                  `json:"path,omitempty"`
}

type ChatReviewRollbackResult struct {
	Review *ChatReviewState `json:"review,omitempty"`
}
