package kitaipu

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

type OptionType int

const (
	OptSubCommand      OptionType = 1
	OptSubCommandGroup OptionType = 2
	OptString          OptionType = 3
	OptInt             OptionType = 4
	OptBool            OptionType = 5
	OptUser            OptionType = 6
	OptChannel         OptionType = 7
	OptRole            OptionType = 8
	OptMentionable     OptionType = 9
	OptNumber          OptionType = 10
	OptAttachment      OptionType = 11
)

type CommandType int

const (
	CmdChatInput CommandType = 1
	CmdUser      CommandType = 2
	CmdMessage   CommandType = 3
)

type InteractionType int

const (
	InteractionPing                           InteractionType = 1
	InteractionApplicationCommand             InteractionType = 2
	InteractionMessageComponent               InteractionType = 3
	InteractionApplicationCommandAutocomplete InteractionType = 4
	InteractionModalSubmit                    InteractionType = 5
)

type CallbackType int

const (
	CallbackPong                                 CallbackType = 1
	CallbackChannelMessageWithSource             CallbackType = 4
	CallbackDeferredChannelMessageWithSource     CallbackType = 5
	CallbackDeferredUpdateMessage                CallbackType = 6
	CallbackUpdateMessage                        CallbackType = 7
	CallbackApplicationCommandAutocompleteResult CallbackType = 8
	CallbackModeal                               CallbackType = 9
)

type option struct {
	Name  string      `json:"name"`
	Type  CommandType `json:"type"`
	Value string      `json:"value"`
}

type subcommand struct {
	Name    string   `json:"name"`
	Options []option `json:"options"`
	Type    string   `json:"type"`
}

type Permission int64

func (p *Permission) UnmarshalJSON(b []byte) (err error) {
	s, err := strconv.Unquote(string(b))
	if err != nil {
		return
	}
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return
	}
	*p = Permission(i)
	return
}

func (p Permission) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.FormatInt(int64(p), 10))
}

func (p Permission) Has(permission Permission) bool {
	return p&permission == permission
}

func (p Permission) Add(permission Permission) Permission {
	return p | permission
}

func (p Permission) Remove(permission Permission) Permission {
	return p &^ permission
}

const (
	PermCreateInstantInvite Permission = 1 << iota
	PermKickMembers
	PermBanMembers
	PermAdministrator
	PermManageChannels
	PermManageGuild
	PermAddReactions
	PermViewAuditLog
	PermPrioritySpeaker
	PermPermStream
	PermViewChannel
	PermSendMessages
	PermSendTTSMessages
	PermManageMessages
	PermEmbedLinks
	PermAttachFiles
	PermReadMessageHistory
	PermMentionEveryone
	PermUseExternalEmojis
	PermViewGuildInsights
	PermConnect
	PermSpeak
	PermMuteMembers
	PermDeafenMembers
	PermMoveMembers
	PermUseVAD
	PermChangeNickname
	PermManageNicknames
	PermManageRoles
	PermManageWebhooks
	PermManageEmojisAndStickers
	PermUseApplicationCommands
	PermRequestToSpeak
	PermManageEvents
	PermManageThreads
	PermCreatePublicThreads
	PermCreatePrivateThreads
	PermUseExternalStickers
	PermSendMessagesInThreads
	PermStartEmbeddedActivities
	PermModerateMembers
)

type user struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Avatar        string `json:"avatar"`
}

type member struct {
	Avatar      string     `json:"avatar"`
	JoinedAt    time.Time  `json:"joined_at"`
	Nick        string     `json:"nick"`
	Permissions Permission `json:"permissions"`
	Roles       []string   `json:"roles"`
	User        user       `json:"user"`
}

type resolvedMember map[string]struct {
	Avatar      string     `json:"avatar"`
	JoinedAt    time.Time  `json:"joined_at"`
	Nick        string     `json:"nick"`
	Permissions Permission `json:"permissions"`
	Roles       []string   `json:"roles"`
}

type resolved struct {
	Members map[string]resolvedMember `json:"members"`
	Users   map[string]user           `json:"users"`
}

type subCommandData struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	SubCommands []subcommand `json:"options"`
	Resolved    resolved     `json:"resolved"`
	Type        CommandType  `json:"type"`
}

type commandData struct {
	ID       string      `json:"id"`
	Name     string      `json:"name"`
	Options  []option    `json:"options"`
	Resolved resolved    `json:"resolved"`
	Type     CommandType `json:"type"`
}

type Command struct {
	ApplicationID string          `json:"application_id"`
	ChannelID     string          `json:"channel_id"`
	Data          commandData     `json:"data"`
	GuildID       string          `json:"guild_id"`
	GuildLocale   string          `json:"guild_locale"`
	ID            string          `json:"id"`
	Locale        string          `json:"locale"`
	Member        member          `json:"member"`
	Token         string          `json:"token"`
	Type          InteractionType `json:"type"`
	Version       int             `json:"version"`
}

type CommandWithSubCommand struct {
	ApplicationID string          `json:"application_id"`
	ChannelID     string          `json:"channel_id"`
	Data          subCommandData  `json:"data"`
	GuildID       string          `json:"guild_id"`
	GuildLocale   string          `json:"guild_locale"`
	ID            string          `json:"id"`
	Locale        string          `json:"locale"`
	Member        member          `json:"member"`
	Token         string          `json:"token"`
	Type          InteractionType `json:"type"`
	Version       int             `json:"version"`
}

type PartialAttachment struct {
	ID          string `json:"id"`
	FileName    string `json:"filename"`
	Description string `json:"description"`
}

type InteractionResponse struct {
	TTS             bool                `json:"tts,omitempty"`
	Content         string              `json:"content,omitempty"`
	Embeds          []Embed             `json:"embeds,omitempty"`
	AllowedMentions *AllowedMentions    `json:"allowed_mentions,omitempty"`
	Flags           MessageFlags        `json:"flags,omitempty"`
	Components      []Component         `json:"components,omitempty"`
	Attachments     []PartialAttachment `json:"attachments,omitempty"`
}

func (r InteractionResponse) Prepare() (responseType string, b []byte, err error) {

	if len(r.Attachments) == 0 {
		r.Attachments = nil
		b, err = json.Marshal(r)
		responseType = "application/json"
		return
	}

	// TODO: make attachments work
	err = errors.New("attachments are not supported yet")
	return

}
