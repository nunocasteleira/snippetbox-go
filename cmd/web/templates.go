package main

import "snippetbox-go.casteleira.xyz/internal/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
