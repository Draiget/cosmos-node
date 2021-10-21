package main

import "time"

type Time struct {
	int64
}

func (t *Time) UnmarshalJSON(b []byte) error {
	tm, err := time.Parse(`"`+time.RFC3339+`"`, string(b))
	if err != nil {
		return err
	}

	*t = Time{tm.Unix()}
	return nil
}

type MetricProtocolVersion struct {
	P2P string `json:"p2p"`
	Block string `json:"block"`
	App string `json:"app"`
}

type MetricNodeOtherInfo struct {
	TxIndex string `json:"tx_index"`
	RpcAddress string `json:"rpc_address"`
}

type MetricNodeInfo struct {
	ProtocolVersion MetricProtocolVersion 	`json:"protocol_version"`
	Id 				string 					`json:"id"`
	ListenAddr 		string 					`json:"listen_addr"`
	Network 		string 					`json:"network"`
	Version 		string 					`json:"version"`
	Channels 		string 					`json:"channels"`
	Moniker 		string 					`json:"moniker"`
	Other 			MetricNodeOtherInfo 	`json:"other"`
}

type MetricSyncInfo struct {
	LatestBlockHash 	string 	`json:"latest_block_hash"`
	LatestAppHash 		string 	`json:"latest_app_hash"`
	LatestBlockHeight 	string 	`json:"latest_block_height"`
	LatestBlockTime 	Time 	`json:"latest_block_time"`
	EarliestBlockHash 	string 	`json:"earliest_block_hash"`
	EarliestAppHash 	string 	`json:"earliest_app_hash"`
	EarliestBlockHeight string 	`json:"earliest_block_height"`
	EarliestBlockTime 	Time 	`json:"earliest_block_time"`
	CatchingUp 			bool 	`json:"catching_up"`
}

type MetricValidatorInfo struct {
	Address string `json:"address"`
	VotingPower string `json:"voting_power"`
}

type MetricStatus struct {
	NodeInfo 		MetricNodeInfo 			`json:"node_info"`
	SyncInfo 		MetricSyncInfo 			`json:"sync_info"`
	ValidatorInfo 	MetricValidatorInfo 	`json:"validator_info"`
}

type MetricPeerConnectionStatus struct {
	Duration string `json:"Duration"`
}

type MetricPeerInfo struct {
	NodeInfo 		MetricNodeInfo 				`json:"node_info"`
	IsOutbound 		bool 						`json:"is_outbound"`
	ConnectionInfo 	MetricPeerConnectionStatus 	`json:"connection_status"`
	RemoteIp 		string 						`json:"remote_ip"`
}

type MetricNetwork struct {
	Listening 		bool 			`json:"listening"`
	NumberOfPeers 	string 			`json:"n_peers"`
	Peers[] 		MetricPeerInfo 	`json:"peers"`
}
