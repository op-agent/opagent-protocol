package op

import "testing"

func TestMetaClone_PreservesStructuredValues(t *testing.T) {
	original := Meta{
		"selectedSkillKeys": []any{"skill-plan"},
		"selectedSkillContext": map[string]any{
			"planFilePath": "/tmp/demo/.agent/plan/thread-1.plan.md",
		},
	}

	cloned := original.Clone()
	keys, ok := cloned["selectedSkillKeys"].([]any)
	if !ok || len(keys) != 1 || keys[0] != "skill-plan" {
		t.Fatalf("cloned selectedSkillKeys = %#v", cloned["selectedSkillKeys"])
	}
	ctx, ok := cloned["selectedSkillContext"].(map[string]any)
	if !ok || ctx["planFilePath"] != "/tmp/demo/.agent/plan/thread-1.plan.md" {
		t.Fatalf("cloned selectedSkillContext = %#v", cloned["selectedSkillContext"])
	}

	keys[0] = "changed"
	ctx["planFilePath"] = "changed"
	origKeys := original["selectedSkillKeys"].([]any)
	origContext := original["selectedSkillContext"].(map[string]any)
	if origKeys[0] != "skill-plan" {
		t.Fatalf("original selectedSkillKeys mutated = %#v", original["selectedSkillKeys"])
	}
	if origContext["planFilePath"] != "/tmp/demo/.agent/plan/thread-1.plan.md" {
		t.Fatalf("original selectedSkillContext mutated = %#v", original["selectedSkillContext"])
	}
}
