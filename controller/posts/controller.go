package posts

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/iancoleman/orderedmap"
	"os"
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
func (c *Controller) Show(ctx context.Context, id string) (story *Story, err error) {
	// Redirect to Landing page?
	fileContent, err := os.ReadFile(fmt.Sprintf("public/stories/%s.json", id))
	if err != nil {
		return nil, err
	}
	var orderedMap orderedmap.OrderedMap
	json.Unmarshal(fileContent, &orderedMap)
	return &Story{orderedMap}, nil

}
