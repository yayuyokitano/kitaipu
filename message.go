package kitaipu

type AllowedMentions struct {
	Parse       []string `json:"parse,omitempty"`
	Users       []string `json:"users,omitempty"`
	Roles       []string `json:"roles,omitempty"`
	RepliedUser bool     `json:"replied_user,omitempty"`
}

type MessageFlags int

func (f MessageFlags) Has(messageFlags MessageFlags) bool {
	return f&messageFlags == messageFlags
}

func (f MessageFlags) Add(messageFlags MessageFlags) MessageFlags {
	return f | messageFlags
}

func (f MessageFlags) Remove(messageFlags MessageFlags) MessageFlags {
	return f &^ messageFlags
}

const (
	MessageCrossposted MessageFlags = 1 << iota
	MessageIsCrosspost
	MessageSuppressEmbeds
	MessageSourceMessageDeleted
	MessageUrgent
	MessageHasThread
	MessageEphemeral
	MessageLoading
	MessageFailedToMentionSomeRolesInThread
)
