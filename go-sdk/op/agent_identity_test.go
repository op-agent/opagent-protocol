package op

import "testing"

func TestBuildKeyIncludesHostID(t *testing.T) {
	key := BuildKey("user123", "devbox-a1b2", string(NodeKindAgent), "file:///tmp/a:1/AGENT.md", EnvCloud)
	want := "user123:devbox-a1b2:agent:file:///tmp/a:1/AGENT.md"
	if key != want {
		t.Fatalf("BuildKey() = %q, want %q", key, want)
	}
}

func TestBuildKeyLocalForcesLocalUID(t *testing.T) {
	key := BuildKey("alice", "devbox-a1b2", string(NodeKindSkill), "file:///tmp/SKILL.md", EnvLocal)
	want := "local:devbox-a1b2:skill:file:///tmp/SKILL.md"
	if key != want {
		t.Fatalf("BuildKey() = %q, want %q", key, want)
	}
}

func TestSplitKeyAndNodeKindFromKey(t *testing.T) {
	key := "user1:host-x9k2:tools:file:///tmp/tools/TOOLS.md"
	uid, hostID, kind, uri, ok := SplitKey(key)
	if !ok {
		t.Fatalf("SplitKey() should parse key: %q", key)
	}
	if uid != "user1" || hostID != "host-x9k2" || kind != "tools" || uri != "file:///tmp/tools/TOOLS.md" {
		t.Fatalf("SplitKey() parsed unexpected values: uid=%q hostID=%q kind=%q uri=%q", uid, hostID, kind, uri)
	}

	nodeKind, ok := NodeKindFromKey(key)
	if !ok {
		t.Fatalf("NodeKindFromKey() should parse key: %q", key)
	}
	if nodeKind != NodeKindTools {
		t.Fatalf("NodeKindFromKey() = %q, want %q", nodeKind, NodeKindTools)
	}
}

func TestBuildNodeIDStableAndKindPrefixed(t *testing.T) {
	id1 := BuildNodeID("user1", "host-z7m3", NodeKindAgent, "file:///tmp/a/.agent/AGENT.md", EnvCloud)
	id2 := BuildNodeID("user1", "host-z7m3", NodeKindAgent, "file:///tmp/a/.agent/AGENT.md", EnvCloud)
	if id1 != id2 {
		t.Fatalf("BuildNodeID() should be stable, got %q and %q", id1, id2)
	}
	if len(id1) != len("agent-xxxx") {
		t.Fatalf("BuildNodeID() should be short, got %q", id1)
	}
	if id1[:6] != "agent-" {
		t.Fatalf("BuildNodeID() should have kind prefix, got %q", id1)
	}
}
