package llm

// NormalizeMessages prepares chat history for OpenAI-compatible APIs.
// Some providers (e.g. OpenRouter/AtlasCloud) reject requests when assistant
// tool-call turns omit content or tool results omit content.
func NormalizeMessages(messages []Message) []Message {
	if len(messages) == 0 {
		return messages
	}
	out := make([]Message, len(messages))
	for i, m := range messages {
		out[i] = normalizeMessage(m)
	}
	return out
}

func normalizeMessage(m Message) Message {
	if m.Content == "" {
		switch m.Role {
		case RoleAssistant:
			if len(m.ToolCalls) > 0 {
				m.Content = ""
			}
		case RoleTool:
			m.Content = "(no output)"
		}
	}

	for j := range m.ToolCalls {
		if m.ToolCalls[j].Type == "" {
			m.ToolCalls[j].Type = "function"
		}
		if m.ToolCalls[j].Function.Arguments == "" {
			m.ToolCalls[j].Function.Arguments = "{}"
		}
	}

	return m
}
