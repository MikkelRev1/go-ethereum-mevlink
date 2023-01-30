package mevlink

import (
	// "fmt"
	// "time"

	// "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/log"
	// "github.com/ethereum/go-ethereum/rlp"
	// mlstreamer "github.com/mevlink/streamer-go"
)

func Stream(eth *eth.Ethereum) {
	go func() {
		// var str = mlstreamer.NewStreamer("id", "key", 1)
		// str.OnTransaction(func(txb []byte, noticed, propagated time.Time) {
		// 	tx := new(types.Transaction)
		// 	err := rlp.DecodeBytes(txb, &tx)
		// 	if err != nil {
		// 		log.Error("[ mevlink-streamer ] error decoding ml tx")
		// 	} else {
		// 		validationErrors := eth.TxPool().AddRemotes([]*types.Transaction{tx})
		// 		if validationErrors[0] == nil {
		// 			log.Info("[ mevlink-streamer ] added tx", "hash", tx.Hash(), "noticed", noticed, "propegated", propagated)
		// 		} else {
		// 			fmt.Println(validationErrors)
		// 			log.Info("[ mevlink-streamer ] benign err adding tx", "hash", tx.Hash(), "err", validationErrors[0])
		// 		}
		// 	}
		// })
		// log.Info("[ mevlink-streamer ] started")
		log.Info("[ mevlink-streamer ] not started; please add credentials and uncomment code")
		// if err := str.Stream(); err != nil {
		// 	log.Info("[ mevlink-streamer ] encountered fatal error: ", err)
		// }
	}()
}
