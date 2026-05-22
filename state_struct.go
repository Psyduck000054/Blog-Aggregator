package main

import (
	"github.com/Psyduck000054/Blog-Aggregator/internal/config"
	"github.com/Psyduck000054/Blog-Aggregator/internal/database"
)

type state struct {
	db_ptr  *database.Queries
	cfg_ptr *config.Config
}
