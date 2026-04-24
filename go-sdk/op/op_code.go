package op

type OpCode string

const (
	// agent
	// Deprecated: thread chat submission should use OpThreadSubmit. Kept as a legacy edge adapter.
	OpAgentCall OpCode = "agent/call"
	// Deprecated: thread chat submission should use OpThreadSubmit. Kept as a legacy edge adapter.
	OpAgentContinue   OpCode = "agent/continue"
	OpAgentLoopCreate OpCode = "agent/loop/create"
	OpPromptGet       OpCode = "prompt/get"
	// OpAgentRoots OpCode = "agents/roots" // list agent roots
	// OpAgentGet   OpCode = "agent/get"
	OpAgentScan OpCode = "agent/scan"

	//node
	// OpNodeScan OpCode = "node/scan"
	OpNodeList OpCode = "node/list"
	// OpNodeCached OpCode = "node/cached"
	// OpNodeCall OpCode = "node/call"

	// OpAgentUpsert OpCode = "agent/upsert"
	// OpLoopCreate  OpCode = "agent/loop/create"

	//host
	SystemStarted OpCode = "system/started"
	// SystemNotify  OpCode = "system/notify"
	// SystemConfigGet OpCode = "system/config/get"
	// HostSecretGet OpCode = "host/secret/get"

	// notify
	NotifyMessage OpCode = "notify/message"

	//config/get
	ConfigGet       OpCode = "config/get"
	ConfigSystemGet OpCode = "config/system/get"

	//thread
	OpChatSessionCreate      OpCode = "chat/session/create"
	OpChatSessionFork        OpCode = "chat/session/fork"
	OpChatSessionMetaGet     OpCode = "chat/session/meta/get"
	OpChatSessionMetaUpdate  OpCode = "chat/session/meta/update"
	OpChatThreadSnapshotGet  OpCode = "chat/thread/snapshot/get"
	OpChatReviewList         OpCode = "chat/review/list"
	OpChatReviewResolve      OpCode = "chat/review/resolve"
	OpChatReviewRollback     OpCode = "chat/review/rollback"
	OpEditorCompletion       OpCode = "editor/completion"
	OpEditorCompletionCancel OpCode = "editor/completion/cancel"
	OpThreadSubmit           OpCode = "thread/submit"
	OpThreadCompact          OpCode = "thread/compact"
	OpThreadInterrupted      OpCode = "thread/interrupted"
	OpThreadElicitReply      OpCode = "thread/elicit_reply"
	OpThreadSteer            OpCode = "thread/steer"
	OpThreadFollowUp         OpCode = "thread/follow_up"
	OpThreadFollowUpPromote  OpCode = "thread/follow_up/promote"
	OpThreadQueueGet         OpCode = "thread/queue/get"
	OpThreadQueueRemove      OpCode = "thread/queue/remove"
	OpThreadActiveList       OpCode = "thread/active/list"
	// OpThreadIDGet OpCode = "threadID/get"
	// OpThreadList        OpCode = "thread/list"
	// OpThreadQuery       OpCode = "thread/query"
	// OpThreadDelete      OpCode = "thread/delete"
	// OpThreadIDUpsert    OpCode = "threadID/upsert"
	// OpThreadUpsert      OpCode = "thread/upsert"

	// //user
	// OpUIDList OpCode = "uid/list" // list all UIDs
	// // user profile / user-agent
	// OpUserProfileGet    OpCode = "user/profile/get"
	// OpUserProfileUpsert OpCode = "user/profile/upsert"
	// OpUserAgentList     OpCode = "user/agent/list"
	// OpUserAgentBind     OpCode = "user/agent/bind"
	// OpUserAgentUnbind   OpCode = "user/agent/unbind"

	// //mcp
	// OpMCPToolCall OpCode = "mcp/tool/call"
	// OpToolCall    OpCode = "tool/call"

	// //elicitation
	// OpElicitCreate OpCode = "elicitation/create"
	// OpElicitUpdate OpCode = "elicitation/update"

	// // skill
	// OpSkillUse OpCode = "skill/use"
	// OpSkillGet OpCode = "skill/get"
)
