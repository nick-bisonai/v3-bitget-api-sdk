package config

import "github.com/bisonai/v3-bitget-api-sdk/pkg/constants"

const (
	BaseUrl = "https://api.bitget.com"
	WsUrl   = "wss://ws.bitget.com/mix/v1/stream"

	TimeoutSecond = 30
	SignType      = constants.SHA256
)
