package models

import (
	"strconv"

	qtcore "github.com/therecipe/qt/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
)

const (
	QName = int(qtcore.Qt__UserRole) + iota + 1
	QAddresses
)

type ModelWallets struct {
	qtcore.QAbstractListModel

	WalletEnv core.WalletEnv
	_ func()                      `constructor:"init"`

	_ map[int]*qtcore.QByteArray  `property:"roles"`
	_ []*ModelAddresses        	  `property:"addresses"`

	_ func()				      `slot:"loadModel"`
	_ func([]*ModelAddresses) 	  `slot:"addAddresses"`
}

func (m *ModelWallets) init() {
	m.SetRoles(map[int]*qtcore.QByteArray{
		Name: 			 qtcore.NewQByteArray2("name", -1),
		QAddresses:		 qtcore.NewQByteArray2("qaddresses", -1),
	})

	m.ConnectRowCount(m.rowCount)
	m.ConnectRoleNames(m.roleNames)
	m.ConnectData(m.data)
	m.ConnectLoadModel(m.loadModel)
	m.ConnectAddAddresses(m.addAddresses)
	altManager := core.LoadAltcoinManager()
	walletsEnvs := make([]core.WalletEnv, 0)
	for _, plug := range altManager.ListRegisteredPlugins() {
		walletsEnvs = append(walletsEnvs, plug.LoadWalletEnvs()...)
	}

	m.WalletEnv = walletsEnvs[0]

	m.loadModel()
}

func (m *ModelWallets) rowCount(*qtcore.QModelIndex) int {
	return len(m.Addresses())
}

func (m *ModelWallets) roleNames() map[int]*qtcore.QByteArray {
	return m.Roles()
}

func (m *ModelWallets) data(index *qtcore.QModelIndex, role int) *qtcore.QVariant {
	if !index.IsValid() {
		return qtcore.NewQVariant()
	}

	if index.Row() >= len(m.Addresses()){
		return qtcore.NewQVariant()
	}

	w := m.Addresses()[index.Row()]

	switch role{
	case QName:
		{
			return qtcore.NewQVariant1(w.Name())
		}
	case QAddresses:
		{
			return qtcore.NewQVariant1(w)
		}
	default:
		{
			return qtcore.NewQVariant()
		}
	}
}

func (m *ModelWallets) insertRows(row int, count int) bool {
	m.BeginInsertRows(qtcore.NewQModelIndex(), row, row + count)
	m.EndInsertRows()
	return true
}

func (m *ModelWallets) loadModel() {
	aModels := make([]*ModelAddresses, 0)
	wallets := m.WalletEnv.GetWalletSet().ListWallets()
	if wallets == nil {
		return
	}
	for wallets.Next() {
		addresses, err := wallets.Value().GetLoadedAddresses()
		if err != nil {
			println(err)
			return
		}
		ma := NewModelAddresses(nil)
		ma.SetName(wallets.Value().GetLabel())
		oModels := make([]*ModelOutputs, 0)

		for addresses.Next() {
			a := addresses.Value()
			outputs := a.GetCryptoAccount().ScanUnspentOutputs()
			mo := NewModelOutputs(nil)
			mo.SetAddress(a.String())
			qOutputs := make([]*QOutput, 0)

			for outputs.Next() {
				to := outputs.Value()
				qo := NewQOutput(nil)
				qo.SetOutputID(to.GetId())
				// TODO: Use correct accuracy here
				accuracy := float64(1000000)
				coins := float64(to.GetCoins("SKY")) / accuracy
				qo.SetAddressSky(coins)
				qo.SetAddressCoinHours(Format(to.GetCoins("")))
				qOutputs = append(qOutputs, qo)
			}
			if len(qOutputs) != 0{
				mo.addOutputs(qOutputs)
				oModels = append(oModels, mo)
			}
		}
		ma.addOutputs(oModels)
		aModels = append(aModels, ma)
	}
	m.addAddresses(aModels)
}

func (m *ModelWallets) addAddresses(ma []*ModelAddresses) {
	m.SetAddresses(ma)
	m.insertRows(len(m.Addresses()), len(ma))
}

func Format(n uint64) string {
    in := strconv.FormatUint(n, 10)
    out := make([]byte, len(in)+(len(in)-2+int(in[0]/'0'))/3)
    if in[0] == '-' {
        in, out[0] = in[1:], '-'
    }

    for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
        out[j] = in[i]
        if i == 0 {
            return string(out)
        }
        if k++; k == 3 {
            j, k = j-1, 0
            out[j] = ','
        }
    }
}