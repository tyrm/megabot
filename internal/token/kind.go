package token

// Kind represents the kind of model to encode a token for
type Kind int64

const (
	// KindUnknown is an unknown token type
	KindUnknown Kind = iota
	// KindUser is a token that represents a user
	KindUser
	// KindGroupMembership is a token that represents a group membership
	KindGroupMembership
	// KindChatbotService is a token that represents a chatbot service
	KindChatbotService
)

func (k Kind) String() string {
	switch k {
	case KindUser:
		return "User"
	case KindGroupMembership:
		return "GroupMembership"
	case KindChatbotService:
		return "ChatbotService"
	default:
		return "unknown"
	}
}
