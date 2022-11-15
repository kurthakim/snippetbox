package main

import "snippetbox.kurt.net/internal/models"

type templateData struct {
	Snippet *models.Snippet
	Snippets []*models.Snippet
}
