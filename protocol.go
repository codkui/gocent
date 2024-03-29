package gocent

import (
	"encoding/json"
	"fmt"
)

// Command represents API command to send.
type Command struct {
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

// Error represents API request error.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %d", e.Message, e.Code)
}

// Reply is a server response to command.
type Reply struct {
	Error  *Error          `json:"error"`
	Result json.RawMessage `json:"result"`
}

// ClientInfo represents information about one client connection to Centrifugo.
// This struct used in messages published by clients, join/leave events, presence data.
type ClientInfo struct {
	User     string          `json:"user"`
	Client   string          `json:"client"`
	ConnInfo json.RawMessage `json:"conn_info,omitempty"`
	ChanInfo json.RawMessage `json:"chan_info,omitempty"`
}

// Publication represents message published into channel.
type Publication struct {
	Offset uint64          `json:"offset"`
	Data   json.RawMessage `json:"data"`
	Info   *ClientInfo     `json:"info"`
}

// NodeInfo contains information and statistics about Centrifugo node.
type NodeInfo struct {
	// UID is a unique id of running node.
	UID string `json:"uid"`
	// Name is a name of node (config defined or generated automatically).
	Name string `json:"name"`
	// Version of Centrifugo node.
	Version string `json:"version"`
	// NumClients is a number of clients connected to node.
	NumClients int `json:"num_clients"`
	// NumUsers is a number of unique users connected to node.
	NumUsers int `json:"num_users"`
	// NumChannels is a number of channels on node.
	NumChannels int `json:"num_channels"`
	// Uptime of node in seconds.
	Uptime int `json:"uptime"`
}

// InfoResult is a result of info command.
type InfoResult struct {
	Nodes []NodeInfo `json:"nodes"`
}

// PublishResult is a result of publish command.
type PublishResult struct {
	Offset uint64 `json:"offset"`
	Epoch  string `json:"epoch"`
}

type PublishResponse struct {
	Error  *Error         `json:"error"`
	Result *PublishResult `json:"result"`
}

// BroadcastResult is a result of broadcast command.
type BroadcastResult struct {
	Responses []PublishResponse `json:"responses"`
}

// PresenceResult is a result of presence command.
type PresenceResult struct {
	Presence map[string]ClientInfo `json:"presence"`
}

// PresenceStatsResult is a result of info command.
type PresenceStatsResult struct {
	NumUsers   int32 `json:"num_users"`
	NumClients int32 `json:"num_clients"`
}

// HistoryResult is a result of history command.
type HistoryResult struct {
	Publications []Publication `json:"publications"`
	Offset       uint64        `json:"offset"`
	Epoch        string        `json:"epoch"`
}

type ChannelInfo struct {
	NumUsers int32 `json:"num_users"`
}

// ChannelsResult is a result of channels command.
type ChannelsResult struct {
	Channels map[string]ChannelInfo `json:"channels"`
}
