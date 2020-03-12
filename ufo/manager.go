package ufo

import (
	"github.com/Assetsadapter/beam-adapter/beam"
	"github.com/blocktree/openwallet/log"
)

type WalletManager struct {
	*beam.WalletManager
}


func NewWalletManager() *WalletManager {
	wm := WalletManager{}
	wm.WalletManager =beam.NewWalletManager()

	wm.Config = beam.NewConfig(Symbol)
	wm.Log = log.NewOWLogger(Symbol)
	return &wm
}

