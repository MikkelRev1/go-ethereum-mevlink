package mevlink

import (
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	mlstreamer "github.com/mevlink/streamer-go"
)

func Stream(eth *eth.Ethereum) {
	go func() {
		var str = mlstreamer.NewStreamer("id", "key", 1)
		str.OnTransaction(func(txb []byte, noticed, propagated time.Time) {
			tx := new(types.Transaction)
			err := rlp.DecodeBytes(txb, &tx)
			if err != nil {
				log.Error("[ mevlink-streamer ] error decoding ml tx")
			} else {
				validationError := eth.TxPool().AddRemotes([]*types.Transaction{tx})
				if validationError == nil {
					 log.Info("[ mevlink-streamer ] added tx", "hash", tx.Hash(), "noticed", noticed, "propegated", propagated)
				}
			}
		})
		str.Stream()
		log.Info("[ mevlink-streamer ] started")
	}()
}
