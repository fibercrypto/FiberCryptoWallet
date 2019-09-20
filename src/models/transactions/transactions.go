package transactions

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/models/address"
	"github.com/therecipe/qt/core"
)

func init() {
	TransactionDetails_QmlRegisterType2("HistoryModels", 1, 0, "QTransactionDetail")
}

const (
	Date = int(core.Qt__UserRole) + 1<<iota
	Status
	Type
	Amount
	HoursTraspassed
	HoursBurned
	TransactionID
	Addresses
	Inputs
	Outputs
)

const (
	TransactionStatusConfirmed = iota
	TransactionStatusPending
	TransactionStatusPreview
)

const (
	TransactionTypeSend = iota
	TransactionTypeReceive
	TransactionTypeInternal
)

type TransactionDetails struct {
	core.QObject

	_ *core.QDateTime      `property:"date"`
	_ int                  `property:"status"`
	_ int                  `property:"type"`
	_ string               `property:"amount"`
	_ string               `property:"hoursTraspassed"`
	_ string               `property:"hoursBurned"`
	_ string               `property:"transactionID"`
	_ *address.AddressList `property:"addresses"`
	_ *address.AddressList `property:"inputs"`
	_ *address.AddressList `property:"outputs"`
}
