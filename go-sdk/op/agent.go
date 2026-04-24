package op

import (
	"context"
	"crypto/sha1"
	"encoding/base32"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/rs/xid"
)

// ---------------------------------------------------------------------------
// OpNode — the universal in-memory representation for agents, skills, tools.
// ---------------------------------------------------------------------------

type NodeKind string

const (
	NodeKindAgent NodeKind = "agent"
	NodeKindSkill NodeKind = "skill"
	NodeKindTools NodeKind = "tools"
)

var Systools = []string{
	"shell",
	"bash",
	"read",
	"write",
	"edit",
}

type OpNode struct {
	ID      string   `json:"id"`                                     // short node id, e.g. agent-ab12
	Key     string   `json:"key"`                                    // uid:hostID:kind:uri
	HostID  string   `json:"hostID,omitempty" mapstructure:"hostID"` // host id
	UID     string   `json:"uid"`                                    // owner/tenant identifier
	OpCodes []OpCode `json:"opCodes,omitempty"`
	Kind    string   `json:"kind"` // agent | skill | tools
	URI     string   `json:"uri"`  // resource locator (file://, cloudos://, ...)
	Cwd     string   `json:"cwd"`  // current working directory
	Tags    []string `json:"tags,omitempty"`
	Run     Run      `json:"run,omitempty"`
	Meta    any      `json:"meta,omitempty"` // AgentMeta | SkillMeta | ToolsMeta
}

// ---------------------------------------------------------------------------
// Identity format: uid:hostID:kind:uri
//
// Examples:
//   local:   local:host-a1b2:agent:file:///example/.opagent/agents/li/.agent/AGENT.md
//   local:   local:host-a1b2:tools:file:///example/.opagent/tools/system-tools/TOOLS.md
//   cloud:   user123:host-c3d4:skill:file:///example/skills/search/SKILL.md
//
// The first three colons delimit uid, hostID and kind.
// Everything after the third colon is the URI (catch-all, may contain colons).
// ---------------------------------------------------------------------------

func BuildKey(uid, hostID, kind, uri string, env string) string {
	if env == EnvLocal {
		uid = "local"
	}
	return strings.TrimSpace(uid) + ":" + strings.TrimSpace(hostID) + ":" + strings.TrimSpace(kind) + ":" + strings.TrimSpace(uri)
}

// ---------------------------------------------------------------------------
// Node ID — short deterministic id with kind prefix.
// ---------------------------------------------------------------------------

const shortNodeIDSuffixLength = 4

func NormalizeNodeKind(kind string) string {
	return strings.ToLower(strings.TrimSpace(kind))
}

func SplitKey(key string) (uid, hostID, kind, uri string, ok bool) {
	key = strings.TrimSpace(key)
	if key == "" {
		return "", "", "", "", false
	}
	parts := strings.SplitN(key, ":", 4) // uid:hostID:kind:uri
	if len(parts) < 4 {
		return "", "", "", "", false
	}
	uid = strings.TrimSpace(parts[0])
	hostID = strings.TrimSpace(parts[1])
	kind = strings.TrimSpace(parts[2])
	uri = strings.TrimSpace(parts[3])
	if uid == "" || hostID == "" || kind == "" || uri == "" {
		return "", "", "", "", false
	}
	return uid, hostID, kind, uri, true
}

func NodeKindFromKey(key string) (NodeKind, bool) {
	_, _, rawKind, _, ok := SplitKey(key)
	if !ok {
		return NodeKind(""), false
	}
	kind := NodeKind(NormalizeNodeKind(rawKind))
	switch kind {
	case NodeKindAgent, NodeKindSkill, NodeKindTools:
		return kind, true
	default:
		return NodeKind(""), false
	}
}

// ComputeNodeID returns a short deterministic id suffix from identity.
func ComputeNodeID(identity string) string {
	sum := sha1.Sum([]byte(strings.TrimSpace(identity)))
	encoded := strings.ToLower(base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(sum[:]))
	if len(encoded) < shortNodeIDSuffixLength {
		return encoded
	}
	return encoded[:shortNodeIDSuffixLength]
}

// BuildNodeID builds id in `kind-xxxx` format from uid/hostID/kind/uri.
func BuildNodeID(uid, hostID string, kind NodeKind, uri string, env string) string {
	normalizedKind := NormalizeNodeKind(string(kind))
	switch NodeKind(kind) {
	case NodeKindAgent, NodeKindSkill, NodeKindTools:
	default:
		normalizedKind = strings.TrimSpace(string(kind))
	}
	identity := BuildKey(strings.TrimSpace(uid), strings.TrimSpace(hostID), normalizedKind, strings.TrimSpace(uri), env)
	return normalizedKind + "-" + ComputeNodeID(identity)
}

func BuildNode(uid, hostID string, kind NodeKind, uri string, env string, tags []string, run Run, opCodes []OpCode, meta any) *OpNode {
	return &OpNode{
		ID:      BuildNodeID(uid, hostID, kind, uri, env),
		Key:     BuildKey(uid, hostID, string(kind), uri, env),
		HostID:  strings.TrimSpace(hostID),
		UID:     uid,
		OpCodes: opCodes,
		Kind:    string(kind),
		URI:     uri,
		Tags:    tags,
		Run:     run,
		Meta:    meta,
	}
}

// ---------------------------------------------------------------------------
// URI helpers
// ---------------------------------------------------------------------------

// PathToURI converts a local path to a file:// URI.
// Optional isDir controls whether a trailing slash is appended.
func PathToURI(path string, isDir ...bool) string {
	// Normalize to forward slashes
	p := strings.ReplaceAll(path, "\\", "/")
	// Handle Windows drive letters: C:/path -> /C:/path
	if len(p) >= 2 && p[1] == ':' {
		p = "/" + p
	}
	withTrailingSlash := len(isDir) > 0 && isDir[0]
	if withTrailingSlash && !strings.HasSuffix(p, "/") {
		p += "/"
	}
	return "file://" + p
}

// URIToPath extracts the local path from a file:// URI.
// Returns empty string if URI is not a file:// URI.
func URIToPath(uri string) string {
	if !strings.HasPrefix(uri, "file://") {
		return ""
	}
	p := strings.TrimPrefix(uri, "file://")
	// Handle Windows: /C:/path -> C:/path
	if len(p) >= 3 && p[0] == '/' && p[2] == ':' {
		p = p[1:]
	}
	// Remove trailing slash for directory paths
	p = strings.TrimSuffix(p, "/")
	return p
}

// URIToDir extracts a local directory path from a file URI.
// For file paths that point to a file, its parent directory is returned.
func URIToDir(uri string) string {
	p := URIToPath(uri)
	if p == "" {
		return ""
	}
	p = strings.TrimSpace(p)
	if p == "" {
		return ""
	}
	if strings.HasSuffix(p, "/") {
		return strings.TrimSuffix(p, "/")
	}
	idx := strings.LastIndex(p, "/")
	if idx <= 0 {
		return ""
	}
	return p[:idx]
}

// ---------------------------------------------------------------------------
// Meta types (no prompt — loaded on-demand from URI)
// ---------------------------------------------------------------------------

type AgentMeta struct {
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	Avatar      string   `json:"avatar,omitempty"`
	MaxToken    int64    `json:"maxToken,omitempty"`
	BindAgentID string   `json:"bindAgentID,omitempty"` // optional bind target agent node ID
	ToolServers []string `json:"toolServers,omitempty"` // tool server OpNode IDs
	SysTools    []string `json:"sysTools,omitempty"`    // system tool names; empty = load all
	Skills      []string `json:"skills,omitempty"`      // skill OpNode IDs
	SubAgents   []string `json:"subAgents,omitempty"`   // agent OpNode IDs
}

type SkillMeta struct {
	Slug        string   `json:"slug"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tags        []string `json:"tags,omitempty"`
}

type ToolUse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	InputSchema any    `json:"inputSchema,omitempty"`
}

type ToolSpec struct {
	ServerID    string `json:"serverID"`
	Name        string `json:"name"`
	Sampling    bool   `json:"sampling,omitempty"`
	Description string `json:"description"`
	InputSchema any    `json:"inputSchema,omitempty"`
}

type SystemBinarySpec struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type ToolsMeta struct {
	Name        string             `json:"name"`
	Description string             `json:"description,omitempty"`
	Tools       []*ToolSpec        `json:"tools,omitempty"`
	SystemBins  []SystemBinarySpec `json:"systemBins,omitempty"`
}

// ---------------------------------------------------------------------------
// Run / Schedule
// ---------------------------------------------------------------------------

type Run struct {
	Command   []string  `json:"command,omitempty"`
	URL       string    `json:"url,omitempty"`
	Lifecycle Lifecycle `json:"lifecycle,omitempty"` // onDemand | daemon | scheduled
	Schedule  Schedule  `json:"schedule,omitempty"`
}

type Lifecycle string

const (
	LifecycleOnDemand  Lifecycle = "onDemand"
	LifecycleDaemon    Lifecycle = "daemon"
	LifecycleScheduled Lifecycle = "scheduled"
)

func (r Run) HasEndpoint() bool {
	return len(r.Command) > 0 || strings.TrimSpace(r.URL) != ""
}

// Validate checks whether run config is internally consistent.
func (r Run) Validate() error {
	if len(r.Command) > 0 && strings.TrimSpace(r.URL) != "" {
		return errors.New("run: command and url are mutually exclusive")
	}
	if !r.Schedule.isEmpty() && r.Lifecycle != LifecycleScheduled {
		return errors.New("run: schedule is only allowed when lifecycle=scheduled")
	}
	switch r.Lifecycle {
	case LifecycleScheduled:
		if r.Schedule.isEmpty() {
			return errors.New("run: lifecycle=scheduled requires schedule config")
		}
		if err := r.Schedule.Validate(); err != nil {
			return fmt.Errorf("run: %w", err)
		}
	case LifecycleDaemon, LifecycleOnDemand, "":
		// valid values
	default:
		return fmt.Errorf("run: unknown lifecycle %q", r.Lifecycle)
	}
	return nil
}

// Schedule defines when a scheduled-lifecycle agent should run.
type Schedule struct {
	Cron  string `json:"cron,omitempty"`
	Every string `json:"every,omitempty"`
	Time  string `json:"time,omitempty"`
}

func (s Schedule) isEmpty() bool {
	return strings.TrimSpace(s.Cron) == "" &&
		strings.TrimSpace(s.Every) == "" &&
		strings.TrimSpace(s.Time) == ""
}

func (s Schedule) Validate() error {
	count := 0
	if strings.TrimSpace(s.Cron) != "" {
		count++
	}
	if strings.TrimSpace(s.Every) != "" {
		count++
	}
	if strings.TrimSpace(s.Time) != "" {
		count++
	}

	switch count {
	case 0:
		return errors.New("schedule requires exactly one of cron, every, or time")
	case 1:
		// valid
	default:
		return errors.New("schedule fields cron, every, and time are mutually exclusive")
	}

	if raw := strings.TrimSpace(s.Every); raw != "" {
		if _, err := time.ParseDuration(raw); err != nil {
			return fmt.Errorf("schedule.every must be a valid duration: %w", err)
		}
	}
	if raw := strings.TrimSpace(s.Cron); raw != "" {
		if len(strings.Fields(raw)) != 5 {
			return fmt.Errorf("schedule.cron must use 5 fields, got %d", len(strings.Fields(raw)))
		}
	}
	if raw := strings.TrimSpace(s.Time); raw != "" {
		if _, err := parseScheduleClock(raw); err != nil {
			return fmt.Errorf("schedule.time must use HH:MM or HH:MM:SS: %w", err)
		}
	}
	return nil
}

func parseScheduleClock(value string) (time.Time, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return time.Time{}, errors.New("empty time")
	}
	for _, layout := range []string{"15:04", "15:04:05"} {
		if parsed, err := time.Parse(layout, value); err == nil {
			return parsed, nil
		}
	}
	return time.Time{}, fmt.Errorf("invalid time %q", value)
}

// ---------------------------------------------------------------------------
// Connection types
// ---------------------------------------------------------------------------

type TransportType string

const (
	Stdio          TransportType = "stdio"
	HttpStreamable TransportType = "http_streamable"
)

type OpNodeParams struct {
	OpCode  OpCode `json:"opCode"`
	Meta    `json:"_meta,omitempty"`
	Content Content `json:"content,omitempty"`
}

func (*OpNodeParams) isParams() {}
func (p *OpNodeParams) UnmarshalJSON(data []byte) error {
	type params OpNodeParams
	var wire struct {
		params
		Content *wireContent `json:"content"`
	}
	if err := json.Unmarshal(data, &wire); err != nil {
		return err
	}
	if wire.Content != nil {
		var err error
		if wire.params.Content, err = contentFromWire(wire.Content, nil); err != nil {
			return err
		}
	}
	*p = OpNodeParams(wire.params)
	return nil
}

// ---------------------------------------------------------------------------
// Agent call / OpAgent protocol types
// ---------------------------------------------------------------------------

type CallAgentHandler func(context.Context, *CallAgentRequest) (*CallAgentResult, error)

type AgentListChangedParams struct {
	Meta `json:"_meta,omitempty"`
}

func (x *AgentListChangedParams) isParams() {}

type OpAgentParams struct {
	OpCode  OpCode `json:"opCode"`
	Meta    `json:"_meta,omitempty"`
	Content Content `json:"content,omitempty"`
}

func (*OpAgentParams) isParams() {}

func (p *OpAgentParams) UnmarshalJSON(data []byte) error {
	type params OpAgentParams
	var wire struct {
		params
		Content *wireContent `json:"content"`
	}
	if err := json.Unmarshal(data, &wire); err != nil {
		return err
	}
	if wire.Content != nil {
		var err error
		if wire.params.Content, err = contentFromWire(wire.Content, nil); err != nil {
			return err
		}
	}
	*p = OpAgentParams(wire.params)
	return nil
}

type CallNodeParams struct {
	Meta    `json:"_meta,omitempty"`
	Content Content `json:"content,omitempty"`
}

func (*CallNodeParams) isParams() {}

func (p *CallNodeParams) UnmarshalJSON(data []byte) error {
	type params CallNodeParams
	var wire struct {
		params
		Content *wireContent `json:"content"`
	}
	if err := json.Unmarshal(data, &wire); err != nil {
		return err
	}
	if wire.Content != nil {
		var err error
		if wire.params.Content, err = contentFromWire(wire.Content, nil); err != nil {
			return err
		}
	}
	*p = CallNodeParams(wire.params)
	return nil
}

type CallAgentParams struct {
	AgentID string `json:"agentID"`
	Meta    `json:"_meta,omitempty"`
	Content Content `json:"content,omitempty"`
}

func (*CallAgentParams) isParams() {}

func (p *CallAgentParams) UnmarshalJSON(data []byte) error {
	type params CallAgentParams
	var wire struct {
		params
		Content *wireContent `json:"content"`
	}
	if err := json.Unmarshal(data, &wire); err != nil {
		return err
	}
	if wire.Content != nil {
		var err error
		if wire.params.Content, err = contentFromWire(wire.Content, nil); err != nil {
			return err
		}
	}
	*p = CallAgentParams(wire.params)
	return nil
}

type OpNodeResult struct {
	OpCode  OpCode `json:"opCode"`
	Meta    `json:"_meta,omitempty"`
	Content Content `json:"content"`
}

func (*OpNodeResult) isResult() {}

func (p *OpNodeResult) UnmarshalJSON(data []byte) error {
	type result OpAgentResult
	var wire struct {
		result
		Content *wireContent `json:"content"`
	}
	if err := json.Unmarshal(data, &wire); err != nil {
		return err
	}
	if wire.Content != nil {
		var err error
		if wire.result.Content, err = contentFromWire(wire.Content, nil); err != nil {
			return err
		}
	}
	*p = OpNodeResult(wire.result)
	return nil
}

type OpAgentResult struct {
	OpCode  OpCode `json:"opCode"`
	Meta    `json:"_meta,omitempty"`
	Content Content `json:"content"`
}

func (*OpAgentResult) isResult() {}
func (r *OpAgentResult) UnmarshalJSON(data []byte) error {
	type result OpAgentResult
	var wire struct {
		result
		Content *wireContent `json:"content"`
	}
	if err := json.Unmarshal(data, &wire); err != nil {
		return err
	}
	if wire.Content != nil {
		var err error
		if wire.result.Content, err = contentFromWire(wire.Content, nil); err != nil {
			return err
		}
	}
	*r = OpAgentResult(wire.result)
	return nil
}

type CallAgentResult struct {
	AgentID string `json:"agentID"`
	Meta    `json:"_meta,omitempty"`
	Content Content `json:"content"`
}

type CallNodeResult struct {
	Meta    `json:"_meta,omitempty"`
	Content Content `json:"content"`
}

func (*CallNodeResult) isResult() {}
func (r *CallNodeResult) UnmarshalJSON(data []byte) error {
	type result CallNodeResult
	var wire struct {
		result
		Content *wireContent `json:"content"`
	}
	if err := json.Unmarshal(data, &wire); err != nil {
		return err
	}
	if wire.Content != nil {
		var err error
		if wire.result.Content, err = contentFromWire(wire.Content, nil); err != nil {
			return err
		}
	}
	*r = CallNodeResult(wire.result)
	return nil
}

func (*CallAgentResult) isResult() {}
func (r *CallAgentResult) UnmarshalJSON(data []byte) error {
	type result CallAgentResult
	var wire struct {
		result
		Content *wireContent `json:"content"`
	}
	if err := json.Unmarshal(data, &wire); err != nil {
		return err
	}
	if wire.Content != nil {
		var err error
		if wire.result.Content, err = contentFromWire(wire.Content, nil); err != nil {
			return err
		}
	}
	*r = CallAgentResult(wire.result)
	return nil
}

type serverAgent struct {
	agent   *AgentMeta
	handler CallAgentHandler
}

// ---------------------------------------------------------------------------
// User task / status (unchanged)
// ---------------------------------------------------------------------------

type UserTask struct {
	UID       string   `bson:"uid" json:"uid"`
	AppID     string   `bson:"appID,omitempty" json:"appID,omitempty"`
	TaskID    string   `bson:"taskID,omitempty" json:"taskID,omitempty"`
	TaskName  string   `bson:"taskName,omitempty" json:"taskName,omitempty"`
	ThreadIDs []string `bson:"threadIDs,omitempty" json:"threadIDs,omitempty"`
}

func GetUserTaskID() string {
	return fmt.Sprintf("ut-%s", xid.New().String())
}

type Status string

const (
	Status_Init       Status = "init"
	Status_Pending    Status = "pending"
	Status_Started    Status = "started"
	Status_InProgress Status = "in_progress"
	Status_Completed  Status = "completed"
	Status_Failed     Status = "failed"
	Status_Running    Status = "running"
	Status_Cancelled  Status = "cancelled"
	Status_Stopped    Status = "stopped"
)
