package main

import "snippet-box.kensonjohnson.com/internal/models"

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}
