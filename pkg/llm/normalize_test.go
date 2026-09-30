package llm

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestNormalizeMessages_AssistantToolCallTurnIncludesContent(t *testing.T) {
	msgs := NormalizeMessages([]Message{
		{
			Role: RoleAssistant,
			ToolCalls: []ToolCall{
				{
					ID:   "call_1",
					Type: "function",
					Function: FunctionCall{
						Name:      "list_dir",
						Arguments: `{"path":"."}`,
					},
				},
			},
		},
		{
			Role:       RoleTool,
			ToolCallID: "call_1",
		},
	})

	raw, err := json.Marshal(msgs)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	s := string(raw)
	if !strings.Contains(s, `"content":""`) {
		t.Fatalf("expected explicit empty content on assistant turn, got %s", s)
	}
	if !strings.Contains(s, `"(no output)"`) {
		t.Fatalf("expected placeholder tool content, got %s", s)
	}
}
