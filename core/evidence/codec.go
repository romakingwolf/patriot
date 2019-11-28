package evidence

import (
	amino "github.com/tendermint/go-amino"
	cryptoAmino "github.com/romakingwolf/patriot/core/crypto/encoding/amino"
	"github.com/romakingwolf/patriot/core/types"
)

var cdc = amino.NewCodec()

func init() {
	RegisterEvidenceMessages(cdc)
	cryptoAmino.RegisterAmino(cdc)
	types.RegisterEvidences(cdc)
}

// For testing purposes only
func RegisterMockEvidences() {
	types.RegisterMockEvidences(cdc)
}
