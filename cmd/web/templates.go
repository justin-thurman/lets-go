package main

import "github.com/justin-thurman/snippetbox/internal/models"

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}
