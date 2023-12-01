package posts

import (
	"context"
	"github.com/iancoleman/orderedmap"
)

// Controller for indices
type Controller struct {
}

// Index struct
type Story struct {
	Story orderedmap.OrderedMap
}

// Index of indices
// GET
func (c *Controller) Index(ctx context.Context) {

}
