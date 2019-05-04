package model

type Contact struct {
	Base
	Channel     `json:"channel"`
	ChannelName `json:"channel_name"`
	Icon        `json:"icon"`
	Qrcode      `json:"qrcode"`
	Link        `json:"link"`
}
