package models

import (
	"github.com/fibercrypto/skywallet-go/src/skywallet"
	"sync"

	"github.com/therecipe/qt/core"
)

type QBridge struct {
	core.QObject
	_      func()                         `constructor:"init"`
	_      func()                         `slot:"onCompleted"`
	_      func()                         `slot:"lock"`
	_      func()                         `slot:"unlock"`
	_      func(message string)           `signal:"getPassword"`
	_      func(title, message string)    `signal:"getBip39Word"`
	_      func(title, message string)    `signal:"getSkyHardwareWalletPin"`
	_      func(title, message string)    `signal:"deviceRequireAction"`
	_      func(title, message string)    `signal:"deviceRequireCancelableAction"`
	_      func(title, message string)    `signal:"deviceRequireConfirmableAction"`
	_      func(string)                   `slot:"setResult"`
	_      func() string                  `slot:"getResult"`
	_ string                              `property:"errMessage"`
	result string
	sem    sync.Mutex
	use    sync.Mutex
}

func (b *QBridge) init() {
	b.ConnectLock(b.lock)
	b.ConnectUnlock(b.unlock)
	b.ConnectSetResult(b.setResult)
	b.ConnectGetResult(b.getResult)
	b.ConnectOnCompleted(b.onCompleted)
}

func (b *QBridge) onCompleted() {
	createSkyHardwareWallet(b)
}

func (b *QBridge) BeginUse() {
	b.use.Lock()
	b.SetErrMessage("")
}

func (b *QBridge) EndUse() {
	b.use.Unlock()
}

func (b *QBridge) lock() {
	b.sem.Lock()
}

func (b *QBridge) setResult(result string) {
	b.result = result
}

func (b *QBridge) getResult() string {
	return b.result
}

func (b *QBridge) getOptionalResult() (string, error) {
	if len(b.ErrMessage()) > 0 {
		logWalletsModel.Warningln(b.ErrMessage())
		return "", skywallet.ErrUserCancelledFromInputReader
	}
	return b.result, nil
}

func (b *QBridge) unlock() {
	b.sem.Unlock()
}
