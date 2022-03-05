package kitaipu

type ComponentType int

const (
	ComponentActionRow  ComponentType = 1
	ComponentButton     ComponentType = 2
	ComponentSelectMenu ComponentType = 3
	ComponentTextInput  ComponentType = 4
)

type ComponentStyle int

// Button styles
const (
	ComponentStylePrimary   ComponentStyle = 1
	ComponentStyleSecondary ComponentStyle = 2
	ComponentStyleSuccess   ComponentStyle = 3
	ComponentStyleDanger    ComponentStyle = 4
	ComponentStyleLink      ComponentStyle = 5
)

// Text styles

const (
	ComponentStyleShort     ComponentStyle = 1
	ComponentStyleParagraph ComponentStyle = 2
)

type PartialEmoji struct {
	Name     string `json:"name"`
	ID       string `json:"id"`
	Animated bool   `json:"animated"`
}

type SelectOptions struct {
	Label       string       `json:"label,omitempty"`
	Value       string       `json:"value,omitempty"`
	Description string       `json:"description,omitempty"`
	Emoji       PartialEmoji `json:"emoji,omitempty"`
	Default     bool         `json:"default,omitempty"`
}

type ButtonComponent struct {
	Type     ComponentType  `json:"type"`
	Style    ComponentStyle `json:"style"`
	Label    string         `json:"label"`
	Emoji    PartialEmoji   `json:"emoji"`
	CustomID string         `json:"custom_id"`
	URL      string         `json:"url"`
	Disabled bool           `json:"disabled"`
}

type SelectComponent struct {
	Type        ComponentType   `json:"type"`
	CustomID    string          `json:"custom_id"`
	Options     []SelectOptions `json:"options"`
	Placeholder string          `json:"placeholder"`
	MinValues   int             `json:"min_values"`
	MaxValues   int             `json:"max_values"`
	Disabled    bool            `json:"disabled"`
}

type TextInputComponent struct {
	Type        ComponentType  `json:"type"`
	CustomID    string         `json:"custom_id"`
	Style       ComponentStyle `json:"style"`
	Label       string         `json:"label"`
	MinLength   int            `json:"min_length"`
	MaxLength   int            `json:"max_length"`
	Required    bool           `json:"required"`
	Value       string         `json:"value"`
	Placeholder string         `json:"placeholder"`
}

type Component struct {
	Type        ComponentType   `json:"type,omitempty"`
	Style       ComponentStyle  `json:"style,omitempty"`
	Label       string          `json:"label,omitempty"`
	Emoji       PartialEmoji    `json:"emoji,omitempty"`
	CustomID    string          `json:"custom_id,omitempty"`
	URL         string          `json:"url,omitempty"`
	Disabled    bool            `json:"disabled,omitempty"`
	Options     []SelectOptions `json:"options,omitempty"`
	Placeholder string          `json:"placeholder,omitempty"`
	MinValues   int             `json:"min_values,omitempty"`
	MaxValues   int             `json:"max_values,omitempty"`
	MinLength   int             `json:"min_length,omitempty"`
	MaxLength   int             `json:"max_length,omitempty"`
	Required    bool            `json:"required,omitempty"`
	Value       string          `json:"value,omitempty"`
}

type ActionRow struct {
	Type       ComponentType `json:"type"`
	Components []Component   `json:"components"`
}
