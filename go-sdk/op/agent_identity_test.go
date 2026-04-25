package op

import "testing"

func TestBuildNodeIdentityIncludesHostID(t *testing.T) {
	identity := BuildNodeIdentity("user123", "devbox-a1b2", string(NodeKindAgent), "file:///tmp/a:1/AGENT.md", EnvCloud)
	want := "user123:devbox-a1b2:agent:file:///tmp/a:1/AGENT.md"
	if identity != want {
		t.Fatalf("BuildNodeIdentity() = %q, want %q", identity, want)
	}
}

func TestBuildNodeIdentityLocalForcesLocalUID(t *testing.T) {
	identity := BuildNodeIdentity("alice", "devbox-a1b2", string(NodeKindSkill), "file:///tmp/SKILL.md", EnvLocal)
	want := "local:devbox-a1b2:skill:file:///tmp/SKILL.md"
	if identity != want {
		t.Fatalf("BuildNodeIdentity() = %q, want %q", identity, want)
	}
}

func TestNodeKindFromID(t *testing.T) {
	nodeKind, ok := NodeKindFromID("tools-c123")
	if !ok {
		t.Fatalf("NodeKindFromID() should parse prefixed id")
	}
	if nodeKind != NodeKindTools {
		t.Fatalf("NodeKindFromID() = %q, want %q", nodeKind, NodeKindTools)
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
