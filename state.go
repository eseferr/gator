package main

import (
	"github.com/eseferr/blog-aggregator/internal/config"
	"github.com/eseferr/blog-aggregator/internal/database"
)

type State struct {
	cfg *config.Config
	db *database.Queries
}

