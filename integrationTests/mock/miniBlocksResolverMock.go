package mock

import (
	"github.com/ElrondNetwork/elrond-go-testing/data/block"
	"github.com/ElrondNetwork/elrond-go-testing/p2p"
)

type MiniBlocksResolverMock struct {
	RequestDataFromHashCalled      func(hash []byte) error
	RequestDataFromHashArrayCalled func(hashes [][]byte) error
	ProcessReceivedMessageCalled   func(message p2p.MessageP2P) error
	GetMiniBlocksCalled            func(hashes [][]byte) (block.MiniBlockSlice, [][]byte)
	GetMiniBlocksFromPoolCalled    func(hashes [][]byte) (block.MiniBlockSlice, [][]byte)
}

func (hrm *MiniBlocksResolverMock) RequestDataFromHash(hash []byte) error {
	return hrm.RequestDataFromHashCalled(hash)
}

func (hrm *MiniBlocksResolverMock) RequestDataFromHashArray(hashes [][]byte) error {
	return hrm.RequestDataFromHashArrayCalled(hashes)
}

func (hrm *MiniBlocksResolverMock) ProcessReceivedMessage(message p2p.MessageP2P, _ func(buffToSend []byte)) error {
	return hrm.ProcessReceivedMessageCalled(message)
}

func (hrm *MiniBlocksResolverMock) GetMiniBlocks(hashes [][]byte) (block.MiniBlockSlice, [][]byte) {
	return hrm.GetMiniBlocksCalled(hashes)
}

func (hrm *MiniBlocksResolverMock) GetMiniBlocksFromPool(hashes [][]byte) (block.MiniBlockSlice, [][]byte) {
	return hrm.GetMiniBlocksFromPoolCalled(hashes)
}

// IsInterfaceNil returns true if there is no value under the interface
func (hrm *MiniBlocksResolverMock) IsInterfaceNil() bool {
	if hrm == nil {
		return true
	}
	return false
}
