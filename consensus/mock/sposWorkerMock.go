package mock

import (
	"github.com/ElrondNetwork/elrond-go-testing/consensus"
	"github.com/ElrondNetwork/elrond-go-testing/data"
	"github.com/ElrondNetwork/elrond-go-testing/p2p"
)

type SposWorkerMock struct {
	AddReceivedMessageCallCalled func(
		messageType consensus.MessageType,
		receivedMessageCall func(cnsDta *consensus.Message) bool,
	)
	RemoveAllReceivedMessagesCallsCalled   func()
	ProcessReceivedMessageCalled           func(message p2p.MessageP2P) error
	SendConsensusMessageCalled             func(cnsDta *consensus.Message) bool
	ExtendCalled                           func(subroundId int)
	GetConsensusStateChangedChannelsCalled func() chan bool
	GetBroadcastBlockCalled                func(data.BodyHandler, data.HeaderHandler) error
	GetBroadcastHeaderCalled               func(data.HeaderHandler) error
	ExecuteStoredMessagesCalled            func()
}

func (sposWorkerMock *SposWorkerMock) AddReceivedMessageCall(messageType consensus.MessageType,
	receivedMessageCall func(cnsDta *consensus.Message) bool) {
	sposWorkerMock.AddReceivedMessageCallCalled(messageType, receivedMessageCall)
}

func (sposWorkerMock *SposWorkerMock) RemoveAllReceivedMessagesCalls() {
	sposWorkerMock.RemoveAllReceivedMessagesCallsCalled()
}

func (sposWorkerMock *SposWorkerMock) ProcessReceivedMessage(message p2p.MessageP2P, _ func(buffToSend []byte)) error {
	return sposWorkerMock.ProcessReceivedMessageCalled(message)
}

func (sposWorkerMock *SposWorkerMock) SendConsensusMessage(cnsDta *consensus.Message) bool {
	return sposWorkerMock.SendConsensusMessageCalled(cnsDta)
}

func (sposWorkerMock *SposWorkerMock) Extend(subroundId int) {
	sposWorkerMock.ExtendCalled(subroundId)
}

func (sposWorkerMock *SposWorkerMock) GetConsensusStateChangedChannel() chan bool {
	return sposWorkerMock.GetConsensusStateChangedChannelsCalled()
}

func (sposWorkerMock *SposWorkerMock) BroadcastBlock(body data.BodyHandler, header data.HeaderHandler) error {
	return sposWorkerMock.GetBroadcastBlockCalled(body, header)
}

func (sposWorkerMock *SposWorkerMock) ExecuteStoredMessages() {
	sposWorkerMock.ExecuteStoredMessagesCalled()
}

// IsInterfaceNil returns true if there is no value under the interface
func (sposWorkerMock *SposWorkerMock) IsInterfaceNil() bool {
	if sposWorkerMock == nil {
		return true
	}
	return false
}
