package functions

import (
	"github.com/Psyduck000054/Blog-Aggregator/internal/config"
	"github.com/Psyduck000054/Blog-Aggregator/internal/database"
)

type State struct {
	Db_ptr  *database.Queries
	Cfg_ptr *config.Config
}
