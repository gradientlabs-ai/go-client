package client

// ParticipantType identifies the type of participant who sent a message or
// is assigned to a conversation.
type ParticipantType string

const (
	// ParticipantTypeCustomer is a  customer/end-user.
	ParticipantTypeCustomer ParticipantType = "Customer"

	// ParticipantTypeHumanAgent is a human support agent.
	ParticipantTypeHumanAgent ParticipantType = "Agent"

	// ParticipantTypeBot is an automation/bot other than the Gradient Labs AI agent.
	ParticipantTypeBot ParticipantType = "Bot"

	// ParticipantTypeAIAgent is the Gradient Labs AI agent.
	ParticipantTypeAIAgent ParticipantType = "AI Agent"
)
