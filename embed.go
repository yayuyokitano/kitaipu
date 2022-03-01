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
	Title       string        `json:"title"`
	Description string        `json:"description"`
	URL         string        `json:"url"`
	Timestamp   time.Time     `json:"timestamp"`
	Color       Color         `json:"color"`
	Footer      EmbedFooter   `json:"footer"`
	Image       EmbedVisual   `json:"image"`
	Thumbnail   EmbedVisual   `json:"thumbnail"`
	Video       EmbedVisual   `json:"video"`
	Provider    EmbedProvider `json:"provider"`
	Author      EmbedAuthor   `json:"author"`
	Fields      []EmbedField  `json:"fields"`
}
