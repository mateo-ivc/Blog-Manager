package posts

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
)

// Controller for indices
type Controller struct {
}

// Index struct
type Story struct {
	Story map[string]interface{}
}

// Index of indices
// GET
func (c *Controller) Show(ctx context.Context, id string) (story *Story, err error) {
	// Redirect to Landing page?

	fileContent, err := os.ReadFile(fmt.Sprintf("public/Stories/%s.json", id))
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	json.Unmarshal(fileContent, &result)
	return &Story{result}, nil

}
