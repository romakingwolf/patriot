package commands

import (
	amino "github.com/tendermint/go-amino"
	cryptoAmino "github.com/romakingwolf/patriot/core/crypto/encoding/amino"
)

var cdc = amino.NewCodec()

func init() {
	cryptoAmino.RegisterAmino(cdc)
}
