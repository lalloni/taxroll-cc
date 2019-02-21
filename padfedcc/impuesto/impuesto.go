package impuesto

import (
	"github.com/lalloni/fabric-chains/collection"
	"github.com/pkg/errors"
	"github.com/s7techlab/cckit/router"
)

type Impuesto struct {
	ID     uint64
	OrgID  string
	Codigo string
	Nombre string
}

func AddGroup(parent *router.Group) {
	c := collection.New([]string{"impuesto"}, collection.JSONMarshaller())
	g := parent.Group("impuesto")
	g.Query("List", list(c, &Impuesto{}))
}

func list(col collection.Collection, t interface{}) router.HandlerFunc {
	return func(c router.Context) (interface{}, error) {
		i, err := col.Accessor(c.Stub()).Iterator()
		if err != nil {
			return nil, errors.Wrap(err, "consultando impuestos")
		}
		imps, err := i.ToArray(t)
		if err != nil {
			return nil, errors.Wrap(err, "obteniendo impuestos")
		}
		return imps, nil
	}
}
