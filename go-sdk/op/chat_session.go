package op

import "encoding/json"

type SessionHeader struct {
	Type              string `json:"type"`
	Version           int    `json:"version"`
	ID                string `json:"id"`
	Timestamp         string `json:"timestamp"`
	AgentKey          string `json:"agentKey"`
	CWD               string `json:"cwd"`
	ChatPath          string `json:"chatPath"`
	FileID            string `json:"fileID,omitempty"`
	Title             string `json:"title"`
	ParentThreadID    string `json:"parentThreadID,omitempty"`
	PlanPath          string `json:"planPath,omitempty"`
	ExecutionPlanPath string `json:"executionPlanPath,omitempty"`
}

type SessionEntryBase struct {
	Type      string  `json:"type"`
	ID        string  `json:"id"`
	ParentID  *string `json:"parentId"`
	Timestamp string  `json:"timestamp"`
}

type SessionElicitationRequestEntry struct {
	SessionEntryBase
	RequestID    string         `json:"requestID"`
	Questions    []QuestionInfo `json:"questions"`
	CurrentIndex int            `json:"currentIndex,omitempty"`
}

type SessionElicitationResultEntry struct {
	SessionEntryBase
	RequestID string     `json:"requestID"`
	Action    string     `json:"action"`
	Answers   [][]string `json:"answers,omitempty"`
	Cancelled bool       `json:"cancelled,omitempty"`
	Answer    string     `json:"answer,omitempty"`
}

type SessionInfoEntry struct {
	SessionEntryBase
	Title             string `json:"title,omitempty"`
	ChatPath          string `json:"chatPath,omitempty"`
	FileID            string `json:"fileID,omitempty"`
	PlanPath          string `json:"planPath,omitempty"`
	ExecutionPlanPath string `json:"executionPlanPath,omitempty"`
}

type SessionCompactionEntry struct {
	SessionEntryBase
	Summary          string `json:"summary"`
	FirstKeptEntryID string `json:"firstKeptEntryId"`
	TokensBefore     int64  `json:"tokensBefore,omitempty"`
}

type ChatSessionCreateParams struct {
	AgentKey string `json:"agentKey"`
	CWD      string `json:"cwd"`
	ChatPath string `json:"chatPath"`
	FileID   string `json:"fileID,omitempty"`
	Title    string `json:"title"`
}

type ChatSessionCreateResult struct {
	ThreadID       string `json:"threadID"`
	FileID         string `json:"fileID,omitempty"`
	Title          string `json:"title"`
	Path           string `json:"path,omitempty"`
	ChatPath       string `json:"chatPath,omitempty"`
	ThreadFilePath string `json:"threadFilePath,omitempty"`
}

type ChatSessionMetaQuery struct {
	ThreadID string `json:"threadID,omitempty"`
	FileID   string `json:"fileID,omitempty"`
	ChatPath string `json:"chatPath,omitempty"`
	AgentKey string `json:"agentKey,omitempty"`
}

type ChatSessionMeta struct {
	ThreadID          string `json:"threadID"`
	FileID            string `json:"fileID,omitempty"`
	AgentKey          string `json:"agentKey"`
	CWD               string `json:"cwd"`
	Path              string `json:"path,omitempty"`
	ChatPath          string `json:"chatPath,omitempty"`
	ThreadFilePath    string `json:"threadFilePath,omitempty"`
	Title             string `json:"title"`
	ParentThreadID    string `json:"parentThreadID,omitempty"`
	PlanPath          string `json:"planPath,omitempty"`
	ExecutionPlanPath string `json:"executionPlanPath,omitempty"`
}

type ChatSessionMetaUpdateParams struct {
	ThreadID          string `json:"threadID,omitempty"`
	FileID            string `json:"fileID,omitempty"`
	ChatPath          string `json:"chatPath,omitempty"`
	Title             string `json:"title,omitempty"`
	PlanPath          string `json:"planPath,omitempty"`
	ExecutionPlanPath string `json:"executionPlanPath,omitempty"`
}

type ChatSessionForkParams struct {
	SourceThreadID    string `json:"sourceThreadID,omitempty"`
	SourceFileID      string `json:"sourceFileID,omitempty"`
	SourceChatPath    string `json:"sourceChatPath,omitempty"`
	AgentKey          string `json:"agentKey,omitempty"`
	CWD               string `json:"cwd,omitempty"`
	FileID            string `json:"fileID,omitempty"`
	ChatPath          string `json:"chatPath,omitempty"`
	Title             string `json:"title"`
	PlanPath          string `json:"planPath,omitempty"`
	ExecutionPlanPath string `json:"executionPlanPath,omitempty"`
}

type TurnResultToolResult struct {
	ToolName        string         `json:"toolName"`
	ArgumentsObject map[string]any `json:"argumentsObject,omitempty"`
	ResultText      string         `json:"resultText"`
	IsError         bool           `json:"isError,omitempty"`
}

type TurnResultPayload struct {
	ThreadID          string                 `json:"threadID"`
	FileID            string                 `json:"fileID,omitempty"`
	TurnID            string                 `json:"turnID"`
	AgentKey          string                 `json:"agentKey"`
	Path              string                 `json:"path,omitempty"`
	ChatPath          string                 `json:"chatPath,omitempty"`
	Title             string                 `json:"title"`
	ParentThreadID    string                 `json:"parentThreadID,omitempty"`
	PlanTurn          bool                   `json:"planTurn,omitempty"`
	UserMessage       Message                `json:"userMessage"`
	AssistantText     string                 `json:"assistantText,omitempty"`
	ReasoningText     string                 `json:"reasoningText,omitempty"`
	ToolResults       []TurnResultToolResult `json:"toolResults,omitempty"`
	CanonicalMessages json.RawMessage        `json:"canonicalMessages,omitempty"`
}
