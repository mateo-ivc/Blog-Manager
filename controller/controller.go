package controller

import (
	context "context"
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
func (c *Controller) Index(ctx context.Context) (stories []*Story, err error) {
	var Stories []*Story
	files, err := os.ReadDir("public/Stories/")

	for _, file := range files {
		fileContent, _ := os.ReadFile(fmt.Sprintf("public/Stories/%s", file.Name()))
		var result map[string]interface{}
		json.Unmarshal(fileContent, &result)
		Stories = append(Stories, &Story{Story: result})
	}
	return Stories, nil
}
