package mock

import (
	"math/big"

	"github.com/ElrondNetwork/elrond-go-testing/data/state"
)

type TemporaryAccountsHandlerMock struct {
	AddTempAccountCalled    func(address []byte, balance *big.Int, nonce uint64)
	CleanTempAccountsCalled func()
	TempAccountCalled       func(address []byte) state.AccountHandler
}

func (tahm *TemporaryAccountsHandlerMock) AddTempAccount(address []byte, balance *big.Int, nonce uint64) {
	if tahm.AddTempAccountCalled == nil {
		return
	}

	tahm.AddTempAccountCalled(address, balance, nonce)
}

func (tahm *TemporaryAccountsHandlerMock) CleanTempAccounts() {
	if tahm.CleanTempAccountsCalled == nil {
		return
	}

	tahm.CleanTempAccountsCalled()
}

func (tahm *TemporaryAccountsHandlerMock) TempAccount(address []byte) state.AccountHandler {
	if tahm.TempAccountCalled == nil {
		return nil
	}

	return tahm.TempAccountCalled(address)
}

// IsInterfaceNil returns true if there is no value under the interface
func (tahm *TemporaryAccountsHandlerMock) IsInterfaceNil() bool {
	if tahm == nil {
		return true
	}
	return false
}
