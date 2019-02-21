package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"

	"gitlab.cloudint.afip.gob.ar/blockchain-team/padfed-chaincode-re/padfedcc"
)

func main() {
	cc := padfedcc.New()
	if err := shim.Start(cc); err != nil {
		cc.Logger().Errorf("error starting padfed chaincode: %s", err)
	}
}
