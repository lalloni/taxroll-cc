package padfedcc

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/s7techlab/cckit/router"

	"gitlab.cloudint.afip.gob.ar/blockchain-team/padfed-chaincode-re/padfedcc/impuesto"
)

func New() shim.Chaincode {
	root := router.New("padfedcc")

	impuesto.AddGroup(root)

	return router.NewChaincode(root)
}
