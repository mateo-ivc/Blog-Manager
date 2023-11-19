package posts

import (
	context "context"
	"fmt"
)

// Controller for indices
type Controller struct {
}

// Index struct
type Posts struct {
	ID int
}

// Index of indices
// GET
func (c *Controller) Index(ctx context.Context) (post *Posts, err error) {
	// Redirect to Landing page?
	fmt.Print()
	a := Posts{
		ID: 1,
	}

	return &a, nil
}
