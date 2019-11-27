package node

import "github.com/ElrondNetwork/elrond-go-testing/node/heartbeat"

func (n *Node) HeartbeatMonitor() *heartbeat.Monitor {
	return n.heartbeatMonitor
}

func (n *Node) HeartbeatSender() *heartbeat.Sender {
	return n.heartbeatSender
}
