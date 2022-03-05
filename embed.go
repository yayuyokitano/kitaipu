package kitaipu

import "time"

type Color int

func (c Color) FromRGB(r, g, b int) Color {
	return Color(r<<16 | g<<8 | b)
}

func (c Color) ToRGB() (r, g, b int) {
	return int(c >> 16), int(c>>8) & 0xff, int(c) & 0xff
}

type EmbedFooter struct {
	Text         string `json:"text"`
	IconURL      string `json:"icon_url"`
	ProxyIconURL string `json:"proxy_icon_url"`
}

type EmbedVisual struct {
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
}

type EmbedProvider struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type EmbedAuthor struct {
	Name         string `json:"name"`
	URL          string `json:"url"`
	IconURL      string `json:"icon_url"`
	ProxyIconURL string `json:"proxy_icon_url"`
}

type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type Embed struct {
	Title       string        `json:"title,omitempty"`
	Description string        `json:"description,omitempty"`
	URL         string        `json:"url,omitempty"`
	Timestamp   time.Time     `json:"timestamp,omitempty"`
	Color       Color         `json:"color,omitempty"`
	Footer      EmbedFooter   `json:"footer,omitempty"`
	Image       EmbedVisual   `json:"image,omitempty"`
	Thumbnail   EmbedVisual   `json:"thumbnail,omitempty"`
	Video       EmbedVisual   `json:"video,omitempty"`
	Provider    EmbedProvider `json:"provider,omitempty"`
	Author      EmbedAuthor   `json:"author,omitempty"`
	Fields      []EmbedField  `json:"fields,omitempty"`
}
